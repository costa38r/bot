package threadcache

import (
	"context"

	"github.com/costa38r/bot/config"
	"github.com/go-redis/redis/v8"
)



func NewRedisClient()(*redis.Client, error){
	cfg := config.LoadConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil,err
	}

	return rdb, nil
}

