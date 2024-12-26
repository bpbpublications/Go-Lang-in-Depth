package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"rest_api_mq/producer/environment"
	"rest_api_mq/producer/models"
	"rest_api_mq/producer/utils"
)

func Example(context *gin.Context) {
	var msg models.Message

	request_id := context.GetString("x-request-id")

	if binderr := context.ShouldBindJSON(&msg); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	connectionString := utils.GetEnvVar("RMQ_URL")

	producer := utils.MessagePublisher{
		environment.EXAMPLE_QUEUE,
		connectionString,
	}

	producer.PublishMessage("text/plain", []byte(msg.Message))

	context.JSON(http.StatusOK, gin.H{
		"response": "Message received from Rest API",
	})

}
