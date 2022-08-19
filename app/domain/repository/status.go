package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// create status
	AddStatus(ctx context.Context, status *object.Status, uID int64) (*object.Status, error)
	// delete status
	DeleteStatus(ctx context.Context, sID int64, accout *object.Account) error
	// get status
	FindByID(ctx context.Context, sID int64) (*object.Status, error)
}
