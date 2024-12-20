package threadcache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
}

func CheckIfThreadExists() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Sem senha por padrão
		DB:       0,  // Usar DB padrão
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatalf("Erro ao configurar o cache: %v", err)
	}

	// Consultar o valor do cache
	val, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		fmt.Println("Chave não encontrada no cache")
	} else if err != nil {
		log.Fatalf("Erro ao consultar o cache: %v", err)
	} else {
		fmt.Printf("Valor da chave 'key': %s\n", val)
	}

}
