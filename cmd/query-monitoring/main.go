package main

import (
	"github.com/nahida05/query-monitoring/config"
	"github.com/nahida05/query-monitoring/internal/service"
	"github.com/nahida05/query-monitoring/internal/storage"
	"github.com/nahida05/query-monitoring/pkg/logger"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logger.New,         // initialize new logger
			config.New,         // initialize config
			storage.New,        // initialize storage interface in order to use db methods
			service.New,        // initialize fiber service
			service.NewHandler, // initialize Handler struct to implement http requests
		),
		fx.Invoke(service.Register),
	).Run()
}
