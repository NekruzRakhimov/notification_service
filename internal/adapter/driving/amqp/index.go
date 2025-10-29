package amqp

import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NekruzRakhimov/notification_service/internal/usecase"
	"github.com/streadway/amqp"
	"log"
)

type ConsumersAMQP struct {
	ProductEvenConsumer *amqp.Queue
	AuthConsumer        *amqp.Queue
	Usecase             *usecase.UseCases
	Channel             *amqp.Channel
}

func NewConsumersAMQP(productEvenConsumer *amqp.Queue,
	authConsumer *amqp.Queue,
	usecase *usecase.UseCases,
	Channel *amqp.Channel) *ConsumersAMQP {
	return &ConsumersAMQP{
		ProductEvenConsumer: productEvenConsumer,
		AuthConsumer:        authConsumer,
		Usecase:             usecase,
		Channel:             Channel,
	}
}

func (c *ConsumersAMQP) Run() {
	go func() {
		msgs, err := c.Channel.Consume(
			c.AuthConsumer.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal("Failed to register a consumer:", err)
		}

		fmt.Println("Waiting for messages. Press CTRL+C to exit.")
		for msg := range msgs {
			var data Message
			if err = json.Unmarshal(msg.Body, &data); err != nil {
				log.Fatal(err)
			}

			err = c.Usecase.Sender.Send(context.Background(), data.Recipient, data.Subject, data.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Received message: %s\n", msg.Body)
		}

	}()

	go func() {
		msgs, err := c.Channel.Consume(
			c.ProductEvenConsumer.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal("Failed to register a consumer:", err)
		}

		fmt.Println("Waiting for messages. Press CTRL+C to exit.")
		for msg := range msgs {
			var data Message
			if err = json.Unmarshal(msg.Body, &data); err != nil {
				log.Fatal(err)
			}

			err = c.Usecase.Sender.Send(context.Background(), data.Recipient, data.Subject, data.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Received message: %s\n", msg.Body)
		}

	}()
}

func InitAMQPConsumer(host string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(host)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	return conn, ch
}

func InitProductEventsQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		"product-events-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	return &queue, nil
}

func InitAuthQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		"auth-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	return &queue, nil
}
