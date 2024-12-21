package threadcache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

//VERIFY IF THE CONTACT NUMBER EXISTIS IN A THREAD CACHE
func CheckValueExists(rdb *redis.Client, key string) (bool, error) {
    ctx := context.Background()
    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return false, nil
    } else if err != nil {
        return false, fmt.Errorf("error to verify value in redis: %w", err)
    }
    // A chave existe
    return val != "", nil
}

