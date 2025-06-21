package main

import (
	"context"
	"fmt"

	"go-redis-demo/internal/config"
	"go-redis-demo/internal/redis"
)

func main() {
	// 获取redis配置
	config.LoadConfig("./internal/config")

	redisConfig := redis.RedisConfig{
		Addr:     config.AppConfig.Redis.Addr,
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	}

	// 创建redis客户端
	redisClient := redis.NewRedisClient(redisConfig)

	result := redisClient.Get(context.Background(), "JOHN")

	fmt.Println(result.Val())
}
