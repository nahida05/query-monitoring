package config

import (
	"github.com/spf13/viper"
)

type db struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type server struct {
	Port string
}

type Config struct {
	DB     db
	Server server
}

func New() *Config {
	viper.AutomaticEnv()
	return &Config{
		DB: db{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
		},
		Server: server{
			Port: viper.GetString("SERVICE_PORT"),
		},
	}
}
