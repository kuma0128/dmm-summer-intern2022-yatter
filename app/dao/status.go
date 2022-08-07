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
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// CreateStatus : statusを作成
func (r *status) AddStatus(ctx context.Context, status *object.Status, uid int64) (*object.Status, error) {
	result, err := r.db.ExecContext(ctx, "INSERT INTO status (account_id, content) VALUES (?, ?)", uid, status.Content)
	if err != nil {
		return status, fmt.Errorf("%w", err)
	}

	status.Sid, err = result.LastInsertId()
	if err != nil {
		return status, fmt.Errorf("%w", err)
	}
	return status, err
}

// delete status
func (r *status) DeleteStatus(ctx context.Context, sid int64, accout *object.Account) error {
	result, err := r.db.ExecContext(ctx, "SELECT id FROM status WHERE id = ?", sid)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	cnt, _ := result.RowsAffected()
	fmt.Printf("%d\n", cnt)
	if cnt == 0 {
		return fmt.Errorf("negative ID doesn't existe")
	}
	_, err = r.db.ExecContext(ctx, "DELETE FROM status WHERE id = ?", sid)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

//get status : statusを取得
func (r *status) FindStatusByID(ctx context.Context, sid int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "SELECT id, account_id, content, create_at FROM status WHERE id = ?", sid).StructScan(entity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}
