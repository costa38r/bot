package threadcache

import (
	"context"
	"fmt"
)


func (c *RedisClient) StoreData(ctx context.Context,client *RedisClient, key string, value string) error {
	err := client.Set(ctx,key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *RedisClient) GetData(ctx context.Context,client *RedisClient, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	fmt.Println(val)
	return val, nil
}