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

//get account : s_idからaccountを取得
func (r *status) FindAccountByID(ctx context.Context, uid int64) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "SELECT id, username, password_hash, display_name, avatar, header, note, create_at FROM account WHERE id = ?", uid).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}

//follow user
func (r *account) FollowAccount(ctx context.Context, uid int64, followedid int64) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO relation (follower_id, followee_id) VALUES (?, ?)", uid, followedid)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return err
}

//get relation
func (r *account) FindRelationByID(ctx context.Context, uid int64, followedid int64) (bool, error) {
	result, err := r.db.ExecContext(ctx, "SELECT follower_id FROM relation WHERE follower_id = ? AND followee_id = ?", uid, followedid)
	if err != nil {
		return false, fmt.Errorf("%w", err)
	}
	cnt, err := result.RowsAffected()
	var following bool
	if cnt == 0 {
		following = false
	} else {
		following = true
	}
	return following, err
}

//get following
func (r *account) FingFollowerByName(ctx context.Context, uid int64, limit int64) ([]*object.Account, error) {
	var entity []*object.Account

	result, err := r.db.QueryxContext(ctx, `SELECT DISTINCT account.id, account.username,
	account.create_at FROM relation LEFT JOIN account ON relation.follower_id = account.id WHERE account.id = ? LIMIT ?`,
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
