package service

import (
	"github.com/nahida05/query-monitoring/internal/storage/model"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	storage Storage
	logger  *logrus.Logger
}

func NewHandler(storage Storage, logger *logrus.Logger) *Handler {
	return &Handler{
		storage: storage,
		logger:  logger,
	}
}

func (h *Handler) Get(ctx *fiber.Ctx) error {

	reqParams := QueryFilter{}
	err := ctx.QueryParser(&reqParams)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(newCustomError(err.Error()))
	}

	err = reqParams.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(newCustomError(err.Error()))
	}

	storageFilter := model.QueryFilter{
		QueryType:       reqParams.Type,
		ExecTimeSorting: reqParams.Sort,
		PageNumber:      reqParams.Page,
		PageLimit:       reqParams.Limit,
	}

	queryResult, err := h.storage.GetList(ctx.Context(), storageFilter)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(newCustomError(internalServerError))
	}

	response := QueryResult{
		Metadata: Metadata{
			Page:       queryResult.Pagination.Page,
			PerPage:    queryResult.Pagination.PerPage,
			PageCount:  queryResult.Pagination.PageCount,
			TotalCount: queryResult.Pagination.TotalCount,
		},
		Payload: make([]Payload, 0, len(queryResult.Queries)),
	}

	for _, query := range queryResult.Queries {
		response.Payload = append(response.Payload, Payload{
			ID:                query.QueryID,
			Statement:         query.Query,
			MaxExecutionTime:  query.MaxExecutionTime,
			MeanExecutionTime: query.MeanExecutionTime,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
