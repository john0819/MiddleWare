package redis

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisOnce     sync.Once
	redisInstance RedisClient
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type redisClient struct {
	client *redis.Client
}

func (r *redisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(config RedisConfig) RedisClient {
	redisOnce.Do(func() {
		redisInstance = &redisClient{
			client: redis.NewClient(&redis.Options{
				Addr:     config.Addr,
				Password: config.Password,
				DB:       config.DB,
			}),
		}
	})

	return redisInstance
}

// todo: redis锁 / 单元测试
