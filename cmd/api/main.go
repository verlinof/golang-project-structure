package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/verlinof/golang-project-structure/configs/app_config"
	"github.com/verlinof/golang-project-structure/configs/db_config"
	"github.com/verlinof/golang-project-structure/db"
	"github.com/verlinof/golang-project-structure/internal/routes"
)

func main() {
	//Init Global Config
	app_config.Config = app_config.LoadConfig()
	db_config.Config = db_config.LoadConfig()

	//Connect Database
	db.ConnectDatabase()

	//Init GIN ENGINE
	gin.SetMode(app_config.Config.GinMode)
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // How long preflight requests can be cached
	}))

	//Init Router
	routes.InitRoute(app)

	//Run Server
	app.Run(":" + app_config.Config.AppPort)
}
