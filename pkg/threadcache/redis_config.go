package threadcache

import (
	"github.com/go-redis/redis/v8"
)

type RedisClientConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(cfg *RedisClientConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
