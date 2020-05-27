package apiserver

import "github.com/alonelegion/golang_rest_api/internal/app/database"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Database *database.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Database: database.NewConfig(),
	}
}
