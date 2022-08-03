package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// create status
	CreateStatus(ctx context.Context, status *object.Status, account *object.Account) error
}
