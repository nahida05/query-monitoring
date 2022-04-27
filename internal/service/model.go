package service

import (
	"fmt"

	"github.com/nahida05/query-monitoring/internal/storage/model"
	"github.com/nahida05/query-monitoring/internal/util"
)

type QueryFilter struct {
	Type  string `json:"type"`
	Sort  string `json:"sort"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

func (f QueryFilter) Validate() error {
	errText := ""
	if f.Type != "" && !util.Exists(f.Type, model.SELECT, model.INSERT, model.UPDATE, model.DELETE) {
		errText = "unsupported value for type field."
	}

	if f.Sort != "" && !util.Exists(f.Sort, model.ASC, model.DESC) {
		errText += "unsupported value for sort field."
	}

	if f.Page < 0 {
		errText += "unsupported value for page field."
	}
	if f.Limit < 0 {
		errText += "unsupported value for limit field."
	}

	if errText != "" {
		return fmt.Errorf(errText)
	}

	return nil

}

type Payload struct {
	ID                int64   `json:"id"`
	Statement         string  `json:"statement"`
	MaxExecutionTime  float64 `json:"max_exec_time"`
	MeanExecutionTime float64 `json:"mean_exec_time"`
}

type Metadata struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}

type QueryResult struct {
	Payload  []Payload `json:"payload"`
	Metadata Metadata  `json:"metadata"`
}
