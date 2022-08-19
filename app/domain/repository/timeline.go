package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	//get timeline
	FindPublicTimelines(ctx context.Context, maxID int64, sinceID int64, limit int64) ([]*object.Status, error)

	FindHomeTimelines(ctx context.Context, uID int64, maxID int64, sinceID int64, limit int64) ([]*object.Status, error)
}
