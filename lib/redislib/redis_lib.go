package redislib

import (
	"IM/common"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
)

func InitClient() {
	client = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})

	pong, err := client.Ping(common.Ctx).Result()
	fmt.Println("初始化Redis....", pong, err)
}

func GetClient() (c *redis.Client) {
	return client
}
