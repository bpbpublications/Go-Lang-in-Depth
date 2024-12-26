package utils

import (
	"github.com/rs/zerolog/log"

	"github.com/streadway/amqp"
)

type MessagePublisher struct {
	Queue            string
	ConnectionString string
}

func (x MessagePublisher) OnError(err error, msg string) {
	if err != nil {
		log.Err(err).Msgf("Publishing message error '%s' queue. Error message: %s", publisher.Queue, msg)
	}
}

func (publisher MessagePublisher) PublishMessage(contentType string, body []byte) {
	conn, error := amqp.Dial(publisher.ConnectionString)
	publisher.OnError(error, "RabbitMQ not connected")
	defer conn.Close()

	channel, err := conn.Channel()
	publisher.OnError(err, "Channel not opened")
	defer channel.Close()

	q, error := channel.QueueDeclare(
		publisher.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	publisher.OnError(error, "Queue Not declared")

	error = channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		})
	publisher.OnError(error, "message not published")
}
