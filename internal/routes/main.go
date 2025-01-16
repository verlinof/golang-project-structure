package routes

import "github.com/gin-gonic/gin"

func InitRoute(app *gin.Engine) {
	api := app.Group("/api")
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
