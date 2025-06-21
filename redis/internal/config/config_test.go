package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPrintRedisConfig(t *testing.T) {
	LoadConfig(".")

	PrintRedisConfig()
}

func TestLoadConfig(t *testing.T) {
	LoadConfig(".")

	assert.Equal(t, AppConfig.Redis.Addr, "localhost:6379")
	assert.Equal(t, AppConfig.Redis.Password, "")
	assert.Equal(t, AppConfig.Redis.DB, 0)
}
