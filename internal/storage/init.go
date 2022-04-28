package storage

import (
	"context"
	"fmt"

	"github.com/nahida05/query-monitoring/config"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(lc fx.Lifecycle, cfg *config.Config, logger *logrus.Logger) (QueryMonitor, error) {
	fmt.Println("test", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name))
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)),
		&gorm.Config{},
	)
	db = db.Debug()

	if err != nil {
		return QueryMonitor{}, fmt.Errorf("unable to create database %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Print("creating pg_stat_statements extension")
			rows, err := db.Raw("SELECT 1 FROM pg_available_extensions WHERE name = ? and installed_version is not null;", "pg_stat_statements").Rows()
			if err != nil {
				return err
			}
			defer rows.Close()

			for rows.Next() {
				logger.Print("pg_stat_statements already exists")
				return nil
			}
			err = db.Exec("CREATE EXTENSION pg_stat_statements;").Error
			if err != nil {
				return err
			}
			logger.Print("creating pg_stat_statements extension")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			sqlConn, err := db.DB()
			if err != nil {
				return err
			}

			err = sqlConn.Close()

			return err
		},
	})

	return QueryMonitor{db: db}, nil
}
