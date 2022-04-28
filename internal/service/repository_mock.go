package service

import (
	"context"
	"github.com/nahida05/query-monitoring/internal/storage/model"
	"github.com/nahida05/query-monitoring/internal/util"
)

type MockRepository struct {
}

func (m MockRepository) GetList(ctx context.Context, filter model.QueryFilter) (model.QueryResult, error) {
	totalCount := 100
	return model.QueryResult{
		Queries: []model.Query{
			{
				QueryID:           123123123123,
				Query:             "select * from users",
				MaxExecutionTime:  0.38499999999999995,
				MeanExecutionTime: 0.38499999999999995,
			},
		},
		Pagination: model.Pagination{
			Page:       filter.PageNumber,
			PerPage:    filter.PageLimit,
			PageCount:  util.PageCount(totalCount, filter.PageLimit, 10),
			TotalCount: totalCount,
		},
	}, nil
}
