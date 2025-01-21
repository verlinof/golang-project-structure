package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	api := app.Group("/api")

	//Dependencies
	// redisManager := pkg_redis.NewRedisManager(redis_config.Config.Host, redis_config.Config.Password, redis_config.Config.Db)
	// validator := pkg_validation.NewXValidator()

	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
