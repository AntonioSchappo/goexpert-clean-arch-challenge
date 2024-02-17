package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/pkg/events"
	"github.com/streadway/amqp"
)

type CreateOrderHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewCreateOrderHandler(rabbitMQChannel *amqp.Channel) *CreateOrderHandler {
	return &CreateOrderHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *CreateOrderHandler) Handle(event events.IEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
