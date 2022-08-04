package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// create status
	CreateStatus(ctx context.Context, status *object.Status, account *object.Account) error
	// get status
	FindByStatusID(ctx context.Context, s_id int64) (*object.Status, error)
	// get account
	FindByAccountID(ctx context.Context, uid int64) (*object.Account, error)
}
