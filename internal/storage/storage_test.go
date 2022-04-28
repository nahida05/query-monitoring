package storage

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/nahida05/query-monitoring/internal/storage/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newGormMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected on sqlmock.New", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock",
		DriverName:           "postgres",
		Conn:                 dbConn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("failed to open gorm v2 db, got error: %v", err)
	}

	return dbConn, db, mock
}

// assume that all filter params are validated by handler
func TestGetList(t *testing.T) {
	predefinedRows := []string{"queryid", "query", "max_exec_time", "mean_exec_time"}
	unexpectedErr := errors.New("unexpected error")

	type args struct {
		ctx    context.Context
		filter model.QueryFilter
	}

	testCases := map[string]struct {
		mock     func(mock sqlmock.Sqlmock)
		args     args
		expected model.QueryResult
		err      error
	}{
		"with valid params": {
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" WHERE lower(query) like $1`)).
					WithArgs("select%").
					WillReturnRows(sqlmock.NewRows(predefinedRows).
						AddRow(20, "select * from users", 4, 111.111).
						AddRow(21, "select * from addressed", 8, 111.111).
						AddRow(22, "select * from services", 6, 111.111))
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" WHERE lower(query) like $1 ORDER BY max_exec_time asc LIMIT 3`)).
					WithArgs("select%").
					WillReturnRows(sqlmock.NewRows(predefinedRows).
						AddRow(20, "select * from users", 4, 111.111).
						AddRow(21, "select * from addressed", 8, 111.111).
						AddRow(22, "select * from services", 6, 111.111))
			},
			args: args{
				filter: model.QueryFilter{
					QueryType:       "select",
					ExecTimeSorting: "asc",
					PageLimit:       3,
				},
			},
			expected: model.QueryResult{
				Queries: []model.Query{
					{
						QueryID:           20,
						Query:             "select * from users",
						MaxExecutionTime:  4,
						MeanExecutionTime: 111.111,
					},
					{
						QueryID:           21,
						Query:             "select * from addressed",
						MaxExecutionTime:  8,
						MeanExecutionTime: 111.111,
					},
					{
						QueryID:           22,
						Query:             "select * from services",
						MaxExecutionTime:  6,
						MeanExecutionTime: 111.111,
					},
				},
				Pagination: model.Pagination{
					Page:       1,
					PerPage:    3,
					PageCount:  1,
					TotalCount: 3,
				},
			},
		},
		"with db error": {
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" WHERE lower(query) like $1`)).
					WithArgs("select%").
					WillReturnError(unexpectedErr)

				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" WHERE lower(query) like $1 ORDER BY max_exec_time asc LIMIT 3`)).
					WithArgs("select%").
					WillReturnRows(sqlmock.NewRows(predefinedRows).
						AddRow(20, "select * from users", 4, 111.111).
						AddRow(21, "select * from addressed", 8, 111.111).
						AddRow(22, "select * from services", 6, 111.111))
			},
			args: args{
				filter: model.QueryFilter{
					QueryType:       "select",
					ExecTimeSorting: "asc",
					PageLimit:       3,
				},
			},
			expected: model.QueryResult{},
			err:      unexpectedErr,
		},
		"with without type and limit": {
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements"`)).
					WillReturnRows(sqlmock.NewRows(predefinedRows).
						AddRow(20, "select * from users", 4, 111.111).
						AddRow(21, "select * from addressed", 8, 111.110).
						AddRow(22, "select * from services", 6, 111.112))
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "pg_stat_statements" ORDER BY max_exec_time desc LIMIT 10`)).
					WillReturnRows(sqlmock.NewRows(predefinedRows).
						AddRow(20, "select * from users", 4, 111.111).
						AddRow(21, "select * from addressed", 8, 111.110).
						AddRow(22, "select * from services", 6, 111.112))
			},
			args: args{
				filter: model.QueryFilter{
					ExecTimeSorting: "desc",
				},
			},
			expected: model.QueryResult{
				Queries: []model.Query{
					{
						QueryID:           20,
						Query:             "select * from users",
						MaxExecutionTime:  4,
						MeanExecutionTime: 111.111,
					},
					{
						QueryID:           21,
						Query:             "select * from addressed",
						MaxExecutionTime:  8,
						MeanExecutionTime: 111.110,
					},
					{
						QueryID:           22,
						Query:             "select * from services",
						MaxExecutionTime:  6,
						MeanExecutionTime: 111.112,
					},
				},
				Pagination: model.Pagination{
					Page:       1,
					PerPage:    10,
					PageCount:  1,
					TotalCount: 3,
				},
			},
		},
	}

	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {

			con, db, mock := newGormMock(t)
			defer con.Close()

			tt.mock(mock)

			got, err := QueryMonitor{db}.GetList(
				tt.args.ctx,
				tt.args.filter,
			)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expected, got)
		})
	}

}
