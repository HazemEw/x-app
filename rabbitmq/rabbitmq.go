package rabbitmq

import (
	"fmt"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

func Connect() (*amqp.Connection, error) {
	if conn != nil {
		return conn, nil
	}

	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	user := os.Getenv("RABBITMQ_USER")
	pass := os.Getenv("RABBITMQ_PASS")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5672"
	}
	if user == "" {
		user = "guest"
	}
	if pass == "" {
		pass = "guest"
	}

	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)

	var err error
	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(rabbitMQURL)
		if err == nil {
			fmt.Println("Successfully connected to RabbitMQ")

			err = DeclareQueue(os.Getenv("RABBITMQ_QUEUE"))
			if err != nil {
				fmt.Println("Failed to declare queue:", err)
			}
			return conn, nil
		}
		fmt.Println("Connection failed. Retrying in 3 seconds...")
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("Failed to connect to RabbitMQ after multiple attempts")
}

func DeclareQueue(queueName string) error {
	channel, err := GetChannel()
	if err != nil {
		return err
	}
	defer channel.Close()

	args := amqp.Table{
		"x-quorum-type": "quorum",
	}
	_, err = channel.QueueDeclare(
		queueName,
		true,
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		args,  // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	fmt.Println("Queue declared:", queueName)
	return nil
}

func GetChannel() (*amqp.Channel, error) {
	if conn == nil {
		_, err := Connect()
		if err != nil {
			return nil, err
		}
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Failed to open RabbitMQ channel: %w", err)
	}

	return channel, nil
}

func CloseConnection() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing RabbitMQ connection:", err)
			return
		}
		fmt.Println("🔌 RabbitMQ connection closed")
	}
}
