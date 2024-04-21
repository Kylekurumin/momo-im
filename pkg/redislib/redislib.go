package redislib

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var client *redis.Client

func NewRedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println("初始化redis:", pong, err)
}

func GetRedisClient() (c *redis.Client) {
	if client == nil {
		NewRedisClient()
	}
	return client
}
