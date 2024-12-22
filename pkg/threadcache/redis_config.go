package threadcache

import (
	"context"

	"github.com/costa38r/bot/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

// NewRedisClient cria e inicializa um novo cliente Redis e retorna uma instância de RedisClient
func NewRedisClient(ctx context.Context) (*RedisClient, error) {

	cfg := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:      cfg.Redis.DB,
	})

	// Teste a conexão com o Redis
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{Client: client}, nil
}
