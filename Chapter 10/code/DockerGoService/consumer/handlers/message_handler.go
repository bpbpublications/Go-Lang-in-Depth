package handlers

import (
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func HandleMessage(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		log.Err(err).Msg("consumer error")
	}
	log.Info().Msgf("Message received from '%s' queue: %s", queue, string(msg.Body))
}
