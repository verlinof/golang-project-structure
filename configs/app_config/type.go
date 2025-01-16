package app_config

type AppConfig struct {
	GinMode      string `env:"GIN_MODE"`
	AppPort      string `env:"APP_PORT"`
	AppUrl       string `env:"APP_URL"`
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
}
