package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"rest_api_mq/producer/middlewares"
	"rest_api_mq/producer/routers"
)

func CreateApp() *gin.Engine {
	log.Info().Msg("service starting")

	app := gin.New()

	app.Use(gin.Recovery())

	app.SetTrustedProxies(nil)

	log.Info().Msg(" cors, request id, request logging middleware added")
	app.Use(middlewares.CORSMiddleware(), middlewares.RequestID(), middlewares.RequestLogger())

	log.Info().Msg("routers setup")
	routers.SetupRouters(app)

	return app
}
