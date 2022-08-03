package dao

import (
	"context"
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
func (r *status) CreateStatus(ctx context.Context, status *object.Status, account *object.Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO status (account_id, content) VALUES (?, ?)", account.ID, status.Content)
	if err != nil {
		return err
	}
	return err
}
