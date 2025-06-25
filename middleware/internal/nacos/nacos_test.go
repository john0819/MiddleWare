package nacos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNcGetConfig(t *testing.T) {
	nc, err := NewNacosClient("127.0.0.1", 8848, "public")
	if err != nil {
		t.Fatalf("Failed to create Nacos client: %v", err)
	}

	content, err := nc.GetConfig("john.config", "DEFAULT_GROUP")
	if err != nil {
		t.Fatalf("Failed to get config: %v", err)
	}

	fmt.Println("content: ", content)
}

func TestNcPusConfig(t *testing.T) {
	// public的namespace是空的
	nc, err := NewNacosClient("127.0.0.1", 8848, "")
	if err != nil {
		t.Fatalf("Failed to create Nacos client: %v", err)
	}

	result, err := nc.PublishConfig("john.config", "DEFAULT_GROUP", "test")

	fmt.Println("result: ", result)

	assert.NoError(t, err)
	assert.True(t, result)
}

func TestNcRegisterInstance(t *testing.T) {
	nc, err := NewNacosClient("127.0.0.1", 8848, "")
	if err != nil {
		t.Fatalf("Failed to create Nacos client: %v", err)
	}

	success, err := nc.RegisterInstance("JohnService", "123.123.123.123", 8080)
	assert.NoError(t, err)
	assert.True(t, success)
}

func TestNcGetService(t *testing.T) {
	nc, err := NewNacosClient("127.0.0.1", 8848, "")
	if err != nil {
		t.Fatalf("Failed to create Nacos client: %v", err)
	}

	serviceInstances, err := nc.GetService("JohnService")
	assert.NoError(t, err)
	assert.NotNil(t, serviceInstances)

	for _, instance := range serviceInstances {
		fmt.Println("instance: ", instance.Ip, instance.Port)
	}
}
