package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func ConsumeMessages(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		os.Getenv("RABBITMQ_QUEUE"), // queue
		"",                          // consumer
		true,                        // auto-ack
		false,                       // exclusive
		false,                       // no-local
		false,                       // no-wait
		nil,                         // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return
	}

	log.Printf("Waiting for messages on queue: %s", os.Getenv("RABBITMQ_QUEUE"))

	for msg := range msgs {
		log.Printf("Received: %s", msg.Body)
	}
}
