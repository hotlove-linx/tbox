package database

import (
	"context"
	"fmt"
	"log"

	"tbox-backend/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis(cfg *config.RedisConfig) error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Redis connected successfully")
	return nil
}

func GetRedis() *redis.Client {
	return RDB
}
