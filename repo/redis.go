package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

const redisPort = 6379

var redisClient *redis.Client

func init() {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)
	if envAddr := os.Getenv("REDIS_ADDR"); envAddr != "" {
		redisAddr = envAddr
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Printf("warning: failed to connect to redis at %s: %v", redisAddr, err)
	}
}

// Set stores raw bytes under the given key in Redis.
func Set(key string, data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return redisClient.Set(ctx, key, data, 0).Err()
}

// Get retrieves raw bytes for the given key.
// Returns nil, nil when the key does not exist.
func Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("redis get %q: %w", key, err)
	}

	return []byte(val), nil
}

// Delete removes the given key from Redis.
func Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return redisClient.Del(ctx, key).Err()
}
