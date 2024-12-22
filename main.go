package main

import (
	"context"

	"github.com/costa38r/bot/config"
	"github.com/costa38r/bot/pkg/threadcache"
)


func main() {
		// Inicializar a configuração
		config.Initialize()
		// Conectar ao Redis
		ctx := context.Background()

		clientRedis,err:=threadcache.NewRedisClient(ctx)
		if err != nil {
			panic(err)
		}

		clientRedis.GetData(ctx,clientRedis,"8399238123")
		
}