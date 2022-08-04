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

func (r *timeline) FindPublicTimelines(ctx context.Context, max_id int64, since_id int64, limit int64) ([]*object.Status, error) {
	var entity []*object.Status
	//var accounts []*object.Account
	//var result *sqlx.Rows
	//max_id && since_id が空白の時、最新のLimit件まで取得
	if max_id == 0 && since_id == 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT status.id, account_id, content, status.create_at, account.id, account.username "account.username", account.create_at FROM status LEFT JOIN account ON status.account_id = account.id  LIMIT ?`, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		//result2, err := r.db.QueryxContext(ctx, "SELECT * FROM account WHERE id IN (SELECT account_id FROM status LIMIT ?)", limit)
		//result2.StructScan(accounts)
		for result.Next() {
			var aaa object.Status
			err = result.StructScan(&aaa)
			//fmt.Printf("%v", aaa)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			entity = append(entity, &aaa)
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
	//if max_id != 0 && since_id != 0 {
	//	result, err := r.db.QueryxContext(ctx, "SELECT * FROM status WHERE id BETWEEN ? AND ?", since_id, max_id)
	//}
	return entity, nil
}
