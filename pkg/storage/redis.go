package storage

import (
	"github.com/Anttoam/golang-htmx-todos/config"
	"github.com/gofiber/storage/redis/v3"
)

func NewRedisClient(cfg *config.Config) *redis.Storage {
	client := redis.New(redis.Config{
		Host: cfg.Redis.Host,
		Port: cfg.Redis.Port,
	})

	return client
}
