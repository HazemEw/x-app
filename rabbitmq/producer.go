package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(queueName string, message string) error {
	channel, err := GetChannel()
	if err != nil {
		return fmt.Errorf("failed to get RabbitMQ channel: %w", err)
	}
	defer channel.Close()

	arg := amqp.Table{
		"x-quorum-type": "quorum",
	}
	_, err = channel.QueueDeclare(
		queueName,
		true,
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		arg,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Publish the message
	err = channel.Publish(
		"",        // Exchange (empty = default)
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent, // ✅ Persistent message
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	fmt.Println("📨 Published message:", message)
	return nil
}
