package app_config

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

var Config *AppConfig

func LoadConfig() *AppConfig {
	cfg := new(AppConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
