package main

import (
	"bytes"
	"fmt"
)

type Logger struct{}

func (logger *Logger) Log(message string) {
	fmt.Println(message)
}

type Message struct {
	logger *Logger
}

func (m *Message) Get(msg string) string {
	m.logger.Log(msg)
	return "Message: " + msg
}

func NewMessage(logger *Logger) *Message {
	return &Message{logger}
}

type Service struct {
	logger  *Logger
	message *Message
}

func (service *Service) GetAllMessages(msgss ...string) string {
	service.logger.Log("\n::All Messages::\n")

	var result bytes.Buffer

	for _, msg := range msgss {
		result.WriteString(service.message.Get(msg))
	}

	return result.String()
}

func NewService(logger *Logger, message *Message) *Service {
	return &Service{logger, message}
}

func main() {
	logger := &Logger{}
	msg := NewMessage(logger)
	ns := NewService(logger, msg)

	msgss := ns.GetAllMessages(
		"message 1",
		"message 2",
	)
	fmt.Println(msgss)
}
