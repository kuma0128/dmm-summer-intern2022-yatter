package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)

	// TODO: Add Other APIs
	// Create user
	AddAccount(ctx context.Context, account *object.Account) error

	// get account
	FindByID(ctx context.Context, uID int64) (*object.Account, error)

	FollowAccount(ctx context.Context, uID int64, followedid int64) error

	UnFollowAccount(ctx context.Context, uID int64, deleteid int64) error

	FindRelationByID(ctx context.Context, uID int64, followedid int64) (bool, error)

	FindFollowingByID(ctx context.Context, uID int64, limit int64) ([]*object.Account, error)

	FindFollowerByID(ctx context.Context, uID int64, maxID int64, sinceID int64, limit int64) ([]*object.Account, error)
}
