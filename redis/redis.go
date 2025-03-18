package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func ConnectRedis() (*redis.Client, error) {
	if client != nil {
		return client, nil
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	fmt.Println("Connected to Redis")
	return client, nil
}

func CloseRedis() {
	if client != nil {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing Redis connection:", err)
		} else {
			fmt.Println("Redis connection closed")
		}
	}
}
