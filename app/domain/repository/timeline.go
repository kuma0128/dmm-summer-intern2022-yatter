package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	//get timeline
	FindPublicTimelines(ctx context.Context, max_id int64, since_id int64, limit int64) ([]*object.Status, error)

	FindHomeTimelines(ctx context.Context, uid int64, max_id int64, since_id int64, limit int64) ([]*object.Status, error)
}
