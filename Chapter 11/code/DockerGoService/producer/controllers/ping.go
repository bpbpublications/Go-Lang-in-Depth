package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
