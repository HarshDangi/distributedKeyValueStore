package redisClient

import (
	"fmt"

	"github.com/harshdangi/distributedKeyValueStore/config"
	"github.com/redis/go-redis/v9"
)

func InitializeClient() *redis.Client {
	port := config.GetEnvParam("REDIS_PORT")
	host := config.GetEnvParam("REDIS_HOST")
	if port == "" || host == "" {
		fmt.Println("Invalid redis address.")
	}
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: config.GetEnvParam("REDIS_PASSWORD"),
		DB:       0,
	})
}
