package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	customRedis "x-app/redis"
)

var ctx = context.Background()

func SetKey(key string, value string, expiration time.Duration) error {
	client, err := customRedis.ConnectRedis()
	if err != nil {
		return fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	err = client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("Failed to set key '%s': %w", key, err)
	}

	fmt.Printf("Key '%s' set in Redis\n", key)
	return nil
}

func GetKey(key string) (string, error) {
	client, err := customRedis.ConnectRedis()
	if err != nil {
		return "", fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	value, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key '%s' not found in Redis", key)
	} else if err != nil {
		return "", fmt.Errorf("Failed to get key '%s': %w", key, err)
	}

	fmt.Printf("Retrieved key '%s' from Redis\n", key)
	return value, nil
}

func DeleteKey(key string) error {
	client, err := customRedis.ConnectRedis()
	if err != nil {
		return fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	err = client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("Failed to delete key '%s': %w", key, err)
	}

	fmt.Printf("🗑️ Deleted key '%s' from Redis\n", key)
	return nil
}

func SetKeyWithTTL(key string, value string, ttl time.Duration) error {
	client, err := customRedis.ConnectRedis()
	if err != nil {
		return fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	err = client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("Failed to set key '%s' with TTL: %w", key, err)
	}

	fmt.Printf("Key '%s' set with TTL of %s\n", key, ttl)
	return nil
}
