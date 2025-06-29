package nacos

import (
	"fmt"
	"log"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// nacos - 服务发现 + 配置中心

const (
	ListenDataId = "john.config"
	DefaultGroup = "DEFAULT_GROUP"
)

var (
	nacosOnce          sync.Once
	nacosClient        NacosInterface
	nacosConfigContent string
)

// 服务中心
type ServiceDiscovery interface {
	RegisterInstance(serviceName string, ip string, port int) (bool, error)
	GetService(serviceName string) ([]ServiceInstance, error)
}

type ServiceInstance struct {
	ServiceName string
	Ip          string
	Port        int
	Weight      int
	Metadata    map[string]string
	Healthy     bool
	Enabled     bool
}

// 配置中心
type ConfigCenter interface {
	PublishConfig(dataId, group, content string) (bool, error)
	GetConfig(dataId, group string) (string, error)
	ListenConfig(dataId, group string, onChange func(content string)) error
}

func OnChange(content string) {
	fmt.Println("nacos config content: ", content)
	nacosConfigContent = content
}

type NacosInterface interface {
	ServiceDiscovery
	ConfigCenter
}

type NacosClient struct {
	configClient config_client.IConfigClient
	namingClient naming_client.INamingClient
}

func initNacosClient(serverAddr string, port uint64, namespace string) (*NacosClient, error) {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: serverAddr,
			Port:   port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)
	if err != nil {
		return nil, err
	}

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)
	if err != nil {
		return nil, err
	}

	return &NacosClient{
		configClient: configClient,
		namingClient: namingClient,
	}, nil
}

// safe gorotuine
// todo: learning safe goroutine
// recover用来捕获当前goroutine的panic，防止程序崩溃 - 终止panic传播
func GoSafe(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
			}
		}()
		fn()
	}()
}

func NewNacosClient(serverAddr string, port uint64, namespace string) (NacosInterface, error) {
	var err error
	nacosOnce.Do(func() {
		nacosClient, err = initNacosClient(serverAddr, port, namespace)
		if err != nil {
			log.Fatalf("Failed to initialize Nacos client: %v", err)
		}

		// 监听配置
		GoSafe(func() {
			err := nacosClient.ListenConfig(ListenDataId, DefaultGroup, OnChange)
			if err != nil {
				log.Printf("Failed to listen config: %v", err)
			}
		})
	})

	return nacosClient, err
}

func (nc *NacosClient) PublishConfig(dataId, group, content string) (bool, error) {
	success, err := nc.configClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content,
	})
	if err != nil {
		return false, err
	}
	return success, nil
}

func (nc *NacosClient) GetConfig(dataId, group string) (string, error) {
	content, err := nc.configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return "", err
	}
	return content, nil
}

func (nc *NacosClient) ListenConfig(dataId, group string, onChange func(content string)) error {
	return nc.configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			onChange(data)
		},
	})
}

func (nc *NacosClient) RegisterInstance(serviceName string, ip string, port int) (bool, error) {
	success, err := nc.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: serviceName,
		Ip:          ip,
		Port:        uint64(port),
	})
	if err != nil {
		return false, err
	}
	return success, nil
}

func (nc *NacosClient) GetService(serviceName string) ([]ServiceInstance, error) {
	instances, err := nc.namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: serviceName,
	})
	if err != nil {
		return nil, err
	}

	serviceInstances := make([]ServiceInstance, len(instances))
	for i, instance := range instances {
		serviceInstances[i] = ServiceInstance{
			ServiceName: serviceName,
			Ip:          instance.Ip,
			Port:        int(instance.Port),
		}
	}
	return serviceInstances, nil
}
