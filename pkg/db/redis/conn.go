package redis

import (
	"avito-banners/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	redisHost := cfg.Redis.Addr
	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: cfg.Redis.Pass,
		DB:       cfg.Redis.DB,
	})

	return client
}
