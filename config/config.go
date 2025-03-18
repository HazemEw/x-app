package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"x-app/api/routes"
	"x-app/database"
	"x-app/rabbitmq"
	"x-app/redis"
)

func LoadConfig() error {
	return godotenv.Load() // Load .env file
}

func GetRabbitMQURL() string {
	return os.Getenv("RABBITMQ_URL")
}

func GetQueueName() string {
	return os.Getenv("QUEUE_NAME")
}

func SetupConfig() {
	database.ConnectDatabase()
	database.RunMigrations()

	r := gin.Default()

	routes.RegisterRoutes(r)

	_, err2 := redis.ConnectRedis()
	if err2 != nil {
		return
	}

	conn, err := rabbitmq.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer rabbitmq.CloseConnection()

	go rabbitmq.ConsumeMessages(conn)

	port := ":" + os.Getenv("PORT")
	log.Println("Server started on", port)
	r.Run(port)
}
