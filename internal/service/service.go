package service

import (
	"context"
	"fmt"

	"github.com/nahida05/query-monitoring/config"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func New(lc fx.Lifecycle, cfg *config.Config, logger *logrus.Logger) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Print("Starting HTTP server.")

			go app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return app.Shutdown()
		},
	})

	return app
}
