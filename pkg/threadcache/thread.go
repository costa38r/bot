package threadcache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func CheckIfThreadExists() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Sem senha por padrão
		DB:       0,  // Usar DB padrão
	})

	// Definir valor no cache com expiração
	err := rdb.Set(ctx, "user123", "John Doe", 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Redis error:", err)
		return
	}

}
