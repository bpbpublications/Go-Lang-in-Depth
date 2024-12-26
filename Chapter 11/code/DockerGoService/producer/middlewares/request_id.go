package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := uuid.New().String()

		c.Set("x-request-id", id)

		c.Header("x-request-id", id)
		c.Next()
	}
}
