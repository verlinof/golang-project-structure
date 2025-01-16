package db_config

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

var Config *DbConfig

func LoadConfig() *DbConfig {
	cfg := new(DbConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
