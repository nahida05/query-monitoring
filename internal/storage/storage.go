package storage

import (
	"context"
	"fmt"
	"github.com/nahida05/query-monitoring/internal/util"
	"strings"

	"github.com/nahida05/query-monitoring/internal/storage/model"

	"gorm.io/gorm"
)

const (
	defaultSorting = "DESC"
	defaultLimit   = 10
	defaultOffset  = 0
)

type queryMonitor struct {
	db *gorm.DB
}

func (q queryMonitor) GetList(ctx context.Context, filter model.QueryFilter) (model.QueryResult, error) {
	query := q.db.WithContext(ctx)

	result := model.QueryResult{}

	//set default filter options
	if filter.PageLimit == 0 {
		filter.PageLimit = defaultLimit
	}

	if filter.PageNumber == 0 {
		filter.PageNumber = defaultOffset
	}

	if filter.ExecTimeSorting == "" {
		filter.ExecTimeSorting = defaultSorting
	}

	// apply filters
	if filter.QueryType != "" {
		query = query.Where("lower(query) like ?", fmt.Sprintf("%s%s", strings.ToLower(filter.QueryType), "%"))
	}

	//get records count for pagination

	filteredData := []model.Query{}
	err := query.Find(&filteredData).Error
	if err != nil {
		return result, nil
	}

	totalCount := len(filteredData)
	result.Pagination = model.Pagination{
		TotalCount: totalCount,
		PageCount:  util.PageCount(totalCount, filter.PageLimit, defaultLimit),
		PerPage:    filter.PageLimit,
		Page:       filter.PageNumber,
	}

	// apply order by, limit and offset
	query = query.Order(fmt.Sprintf("max_exec_time %s", filter.ExecTimeSorting))

	err = query.Limit(filter.PageLimit).
		Offset((filter.PageNumber - 1) * filter.PageLimit).
		Find(&result.Queries).Error

	return result, err
}
