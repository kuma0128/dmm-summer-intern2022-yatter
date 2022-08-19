package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) FindPublicTimelines(ctx context.Context, maxID int64, sinceID int64, limit int64) ([]*object.Status, error) {
	var entity []*object.Status
	//var accounts []*object.Account
	//var result *sqlx.Rows
	//max_id && since_id が空白の時、最新のLimit件まで取得
	if maxID == 0 && sinceID == 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT status.id, status.account_id, status.content,
		 status.create_at, account.id "account.id", account.username "account.username",
		 account.create_at "account.create_at" FROM status LEFT JOIN account ON status.account_id = account.id  LIMIT ?`, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		// fmt.Printf("%d\n", max_id)
		// fmt.Printf("%d\n", since_id)
		// fmt.Printf("%d\n", limit)
		//result2, err := r.db.QueryxContext(ctx, "SELECT * FROM account WHERE id IN (SELECT account_id FROM status LIMIT ?)", limit)
		//result2.StructScan(accounts)
		for result.Next() {
			var tmp object.Status
			err = result.StructScan(&tmp)
			//fmt.Printf("%v\n", tmp)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			entity = append(entity, &tmp)
		}
		//err = result.StructScan(entity)
		//var account_ids []int64
		// set account information
		// for i, _ := range entity {
		// 	account_ids[i] = entity[i].AccountID
		// }
		return entity, err
	}
	// since_id <= x <= max_id
	if maxID != 0 && sinceID != 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT status.id, status.account_id, status.content,
		 status.create_at, account.id "account.id", account.username "account.username",
		 account.create_at "account.create_at" FROM status LEFT JOIN account ON status.account_id = account.id 
		 WHERE status.id BETWEEN ? AND ? LIMIT ?`, sinceID, maxID, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		for result.Next() {
			var tmp object.Status
			err = result.StructScan(&tmp)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			entity = append(entity, &tmp)
		}
		return entity, err
	}

	return entity, nil
}

func (r *timeline) FindHomeTimelines(ctx context.Context, uID int64, maxID int64, sinceID int64, limit int64) ([]*object.Status, error) {
	var entity []*object.Status

	if maxID == 0 && sinceID == 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT DISTINCT status.id, status.account_id, status.content,
		status.create_at, account.id "account.id", account.username "account.username",
		account.create_at "account.create_at" FROM status 
		LEFT JOIN account ON status.account_id = account.id 
		INNER JOIN relation ON  status.account_id IN(SELECT followee_id FROM relation WHERE follower_id = ?) LIMIT ?`, uID, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		for result.Next() {
			var tmp object.Status
			err = result.StructScan(&tmp)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			entity = append(entity, &tmp)
		}
		return entity, err
	}

	if maxID != 0 && sinceID != 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT DISTINCT status.id, status.account_id, status.content,
		status.create_at, account.id "account.id", account.username "account.username",
		account.create_at "account.create_at" FROM status 
		LEFT JOIN account ON status.account_id = account.id 
		INNER JOIN relation ON  status.account_id IN(SELECT followee_id FROM relation WHERE follower_id = ?)
		WHERE status.id BETWEEN ? AND ? LIMIT ?`, uID, sinceID, maxID, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		for result.Next() {
			var tmp object.Status
			err = result.StructScan(&tmp)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			entity = append(entity, &tmp)
		}
		return entity, err
	}

	return entity, nil
}
