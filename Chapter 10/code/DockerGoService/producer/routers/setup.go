package routers

import (
	"github.com/gin-gonic/gin"
	"rest_api_mq/producer/controllers"
)

func CreateRouters(engine *gin.Engine) {
	version1 := engine.Group("/v1")
	{
		version1.GET("/ping", controllers.Ping)
		version1.POST("/publish/example", controllers.Example)
	}

}
