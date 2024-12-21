package threadcache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func CheckIfThreadExists(ctx context.Context, cfg *RedisClientConfig, key, value string) {
	rdb := NewRedisClient(cfg)

	// Consulta o valor no cache
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("Chave '%s' não encontrada no cache\n", key)
	} else if err != nil {
		log.Fatalf("Erro ao consultar o cache: %v", err)
	} else {
		fmt.Printf("Valor da chave '%s': %s\n", key, val)
	}

}
