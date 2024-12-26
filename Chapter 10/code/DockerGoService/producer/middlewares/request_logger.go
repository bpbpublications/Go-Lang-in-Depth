package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RequestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		request_id := context.GetString("x-request-id")
		client_ip := context.ClientIP()
		user_agent := context.Request.UserAgent()
		method := context.Request.Method
		path := context.Request.URL.Path

		timeNow := time.Now()

		context.Next()

		latency := float32(time.Since(timeNow).Seconds())

		status := context.Writer.Status()
		log.Info().Str("request_id", request_id).Str("client_ip", client_ip).
			Str("user_agent", user_agent).Str("method", method).Str("path", path).
			Float32("latency", latency).Int("status", status).Msg("")

	}
}
