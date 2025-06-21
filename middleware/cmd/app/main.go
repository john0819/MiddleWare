package main

import (
	"context"
	"fmt"

	"go-redis-demo/internal/config"
	"go-redis-demo/internal/freecache"
	"go-redis-demo/internal/pkg/utils"
	"go-redis-demo/internal/redis"
	"go-redis-demo/internal/user"

	"github.com/zeromicro/go-zero/zrpc"
)

func init() {
	freecache.InitCacheManager()
}

func StartService() {
	userConfig := config.ServiceConfigInstance.User
	user.RunUserService(zrpc.RpcServerConf{
		ListenOn: fmt.Sprintf("%s:%d", userConfig.Host, userConfig.Port),
	})
}

func main() {
	// 获取redis配置
	config.LoadConfig("./internal/config")

	// 获取服务配置
	config.LoadServiceConfig("./etc")

	redisConfig := redis.RedisConfig{
		Addr:     config.AppConfig.Redis.Addr,
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	}

	// 创建redis客户端
	redisClient := redis.NewRedisClient(redisConfig)

	result := redisClient.Get(context.Background(), "JOHN")

	fmt.Println(result.Val())

	// freecache manager
	cacheManager := freecache.GetCacheManager()
	cacheSize := utils.FreeCacheSize(config.AppConfig.FreeCache.JohnCacheSize)
	johnCache := cacheManager.GetOrCreateCache("john", cacheSize)
	johnCache.Set([]byte("john"), []byte("wong cache"), 0)
	cacheResult, err := johnCache.Get([]byte("john"))
	if err != nil {
		fmt.Println("get john cache error", err)
	}
	fmt.Println(string(cacheResult))

	StartService()
}
