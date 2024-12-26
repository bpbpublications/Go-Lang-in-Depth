package main

import (
	"rest_api_mq/consumer/environment"
	"rest_api_mq/consumer/handlers"
	"rest_api_mq/consumer/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")

	messageQueue := utils.MessageConsumer{
		enviornment.EXAMPLE_QUEUE,
		connectionString,
		handlers.HandleMessaage,
	}

	forever := make(chan bool)

	go messageQueue.Consume()

	<-forever
}
