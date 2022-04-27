package model

const (
	StatementsTableName = "pg_stat_statements"
)

const (
	SELECT = "select"
	UPDATE = "update"
	INSERT = "insert"
	DELETE = "delete"

	ASC  = "asc"
	DESC = "desc"
)

type QueryFilter struct {
	QueryType       string
	ExecTimeSorting string
	PageNumber      int
	PageLimit       int
}

type Query struct {
	QueryID           int64   `gorm:"column:queryid"`
	Query             string  `gorm:"column:query"`
	MaxExecutionTime  float64 `gorm:"column:max_exec_time"`
	MeanExecutionTime float64 `gorm:"column:mean_exec_time"`
}

type Pagination struct {
	Page       int
	PerPage    int
	PageCount  int
	TotalCount int
}

type QueryResult struct {
	Queries    []Query
	Pagination Pagination
}

func (g Query) TableName() string {
	return StatementsTableName
}
