package main

import (
	"github.com/verlinof/golang-project-structure/configs/app_config"
	"github.com/verlinof/golang-project-structure/configs/db_config"
)

func main() {
	//Init Global Config
	app_config.Config = app_config.LoadConfig()
	db_config.Config = db_config.LoadConfig()
}
