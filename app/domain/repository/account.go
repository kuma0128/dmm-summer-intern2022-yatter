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

	FollowAccount(ctx context.Context, uid int64, followedid int64) error

	FindRelationByID(ctx context.Context, uid int64, followedid int64) (bool, error)

	FindFollowingByName(ctx context.Context, uid int64, limit int64) ([]*object.Account, error)

	FindFollowerByName(ctx context.Context, uid int64, max_id int64, since_id int64, limit int64) ([]*object.Account, error)
}
