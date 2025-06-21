package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 是我们整个应用配置的根结构体
type Config struct {
	Redis     RedisConfig     `mapstructure:"redis"`
	FreeCache FreeCacheConfig `mapstructure:"freecache"`
}

// RedisConfig 包含了所有与 Redis 相关的配置项
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type FreeCacheConfig struct {
	JohnCacheSize int `mapstructure:"johnCacheSize"`
}

// AppConfig 是一个全局变量，将在加载后持有所有配置信息
var AppConfig Config

// 服务配置
type ServiceConfig struct {
	User UserConfig `mapstructure:"user"`
}

type UserConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var ServiceConfigInstance ServiceConfig

func PrintRedisConfig() {
	fmt.Println("关于redis配置的细节")
	fmt.Println("redis addr: ", AppConfig.Redis.Addr)
	fmt.Println("redis password: ", AppConfig.Redis.Password)
	fmt.Println("redis db: ", AppConfig.Redis.DB)
}

// 服务配置
func LoadServiceConfig(path string) {
	// 重置 viper 状态
	viper.Reset()

	viper.SetConfigName("user")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("致命错误：无法读取配置文件: %w", err))
	}

	if err := viper.Unmarshal(&ServiceConfigInstance); err != nil {
		panic(fmt.Errorf("致命错误：无法解码服务配置: %w", err))
	}
}

// 通用配置
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
