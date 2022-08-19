package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

func (r *account) AddAccount(ctx context.Context, account *object.Account) error {
	//entity := new(object.Account)
	_, err := r.db.ExecContext(ctx, "INSERT INTO account (username, password_hash) VALUES (?, ?)", account.Username, account.PasswordHash)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return err
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// get account : s_idからaccountを取得
func (r *account) FindByID(ctx context.Context, uID int64) (*object.Account, error) {
	entity := new(object.Account)
	if err := r.db.QueryRowxContext(ctx, "SELECT id, username, password_hash, display_name, avatar, header, note, create_at FROM account WHERE id = ?", uID).StructScan(entity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}

// follow user
func (r *account) FollowAccount(ctx context.Context, uID int64, followedID int64) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO relation (follower_id, followee_id) VALUES (?, ?)", uID, followedID); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// unfollow user
func (r *account) UnFollowAccount(ctx context.Context, uid int64, deleteid int64) error {
	var err error
	if _, err = r.db.ExecContext(ctx, "DELETE FROM relation WHERE follower_id = ? AND followee_id = ?", uid, deleteid); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// get relation
func (r *account) FindRelationByID(ctx context.Context, uid int64, followedid int64) (bool, error) {
	result, err := r.db.QueryxContext(ctx, "SELECT follower_id FROM relation WHERE follower_id = ? AND followee_id = ?", uid, followedid)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	//fmt.Printf("%d\n", cnt)
	var following bool
	if result == nil {
		following = false
	} else {
		following = true
	}
	return following, err
}

// get following
func (r *account) FindFollowingByID(ctx context.Context, uid int64, limit int64) ([]*object.Account, error) {
	var entity []*object.Account

	result, err := r.db.QueryxContext(ctx, `SELECT DISTINCT account.id, account.username,
	account.create_at FROM relation LEFT JOIN account ON relation.follower_id = account.id AND relation.follower_id = ? LIMIT ?`,
		uid, limit)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	for result.Next() {
		var tmp object.Account
		err = result.StructScan(&tmp)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		entity = append(entity, &tmp)
	}

	return entity, err
}

func (r *account) FindFollowerByID(ctx context.Context, uid int64, max_id int64, since_id int64, limit int64) ([]*object.Account, error) {
	var entity []*object.Account
	if max_id == 0 && since_id == 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT DISTINCT account.id, account.username,
		account.create_at FROM relation LEFT JOIN account ON relation.followee_id = account.id AND relation.follower_id = ? LIMIT ?`,
			uid, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		for result.Next() {
			var tmp object.Account
			err = result.StructScan(&tmp)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			entity = append(entity, &tmp)
		}
		return entity, err
	}

	if max_id != 0 && since_id != 0 {
		result, err := r.db.QueryxContext(ctx, `SELECT DISTINCT account.id, account.username,
		account.create_at FROM relation LEFT JOIN account ON relation.followee_id = account.id AND relation.follower_id = ? WHERE account.id BETWEEN ? AND ? LIMIT ?`,
			uid, max_id, since_id, limit)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		for result.Next() {
			var tmp object.Account
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
