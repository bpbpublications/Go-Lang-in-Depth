package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type MessageConsumer struct {
	Queue            string
	ConnectionString string
	MsgHandler       func(queue string, msg amqp.Delivery, err error)
}

func (consumer MessageConsumer) OnError(errors error, msg string) {
	if errors != nil {
		consumer.MsgHandler(consumer.Queue, amqp.Delivery{}, errors)
	}
}

func (consumer MessageConsumer) Consume() {
	conn, err := amqp.Dial(consumer.ConnectionString)
	consumer.OnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	consumer.OnError(err, "Failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		consumer.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	consumer.OnError(err, "Failed to declare a queue")

	msgs, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	consumer.OnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for delivery := range msgs {
			consumer.MsgHandler(consumer.Queue, delivery, nil)
		}
	}()
	log.Info().Msgf("Started listening- messages from '%s' queue", consumer.Queue)
	<-forever
}
