package repository

import (
	"context"
	"database/sql"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// create status
	AddStatus(ctx context.Context, status *object.Status, account *object.Account) (*object.Status, sql.Result, error)
	// delete status
	DeleteStatus(ctx context.Context, sid int64, accout *object.Account) error
	// get status
	FindStatusByID(ctx context.Context, sid int64) (*object.Status, error)
	// get account
	FindAccountByID(ctx context.Context, uid int64) (*object.Account, error)
}
