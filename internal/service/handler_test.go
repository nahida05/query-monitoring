package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/magiconair/properties/assert"
	"github.com/nahida05/query-monitoring/internal/storage/repository"
	"github.com/nahida05/query-monitoring/pkg/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Get(t *testing.T) {

	h := NewHandler(
		repository.Mock{},
		logger.New(),
	)

	tests := []struct {
		name               string
		route              string
		expectedStatusCode int
	}{
		{
			name:               "with wrong type param",
			route:              "/queries?type=create&sort=asc&limit=10&page=2",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "with wrong limit param",
			route:              "/queries?type=select&sort=asc&limit=abs&page=2",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "with wrong sort param",
			route:              "/queries?type=select&sort=asc-desc&limit=10&page=2",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "with success",
			route:              "/queries?type=select&sort=asc&limit=10&page=2",
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/queries", h.Get)
			resp, err := app.Test(httptest.NewRequest(
				http.MethodGet,
				tt.route,
				nil,
			))

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}
