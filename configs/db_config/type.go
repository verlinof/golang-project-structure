package db_config

type DbConfig struct {
	Driver     string `env:"DB_DRIVER"`
	Host       string `env:"DB_HOST"`
	Port       string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASS"`
	DbName     string `env:"DB_NAME"`
}
