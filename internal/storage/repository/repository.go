package repository

import (
	"context"
	"github.com/nahida05/query-monitoring/internal/storage/model"
)

type Storage interface {
	GetList(ctx context.Context, filter model.QueryFilter) (model.QueryResult, error)
}
