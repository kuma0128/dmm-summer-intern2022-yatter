package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	//get timeline
	FindPublicTimelines(ctx context.Context, max_id int64, since_id int64, limit int32, statuses []object.Status) ([]object.Status, error)
}
