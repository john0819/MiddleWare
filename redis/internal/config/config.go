package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 是我们整个应用配置的根结构体
type Config struct {
	Redis RedisConfig `mapstructure:"redis"`
}

// RedisConfig 包含了所有与 Redis 相关的配置项
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// AppConfig 是一个全局变量，将在加载后持有所有配置信息
var AppConfig Config

func PrintRedisConfig() {
	fmt.Println("关于redis配置的细节")
	fmt.Println("redis addr: ", AppConfig.Redis.Addr)
	fmt.Println("redis password: ", AppConfig.Redis.Password)
	fmt.Println("redis db: ", AppConfig.Redis.DB)
}

func LoadConfig(path string) {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("致命错误：无法读取配置文件: %w", err))
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("致命错误：无法解码配置: %w", err))
	}
}
