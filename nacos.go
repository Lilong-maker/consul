package consul

import (
	"fmt"

	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Mysql
	Redis
	//Nacos
}
type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

var (
	Gen *AppConfig
)

func NacosInit() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "115.190.43.83",
			Port:   uint64(8848),
		},
	}

	nacosClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		fmt.Printf("创建 Nacos 客户端失败: %v\n", err)
		return
	}
	configContent, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: "order",
		Group:  "dev",
	})
	if err != nil {
		fmt.Printf("从 Nacos 获取配置失败: %v\n", err)
		return
	}
	fmt.Println(configContent)
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(strings.NewReader(configContent))
	if err != nil {
		fmt.Printf("解析 Nacos 配置失败: %v\n", err)
		return
	}
	err = viper.Unmarshal(&Gen)
	if err != nil {
		fmt.Printf("反序列化 Nacos 配置失败: %v\n", err)
		return
	}

	fmt.Println("Nacos 配置读取成功")
}
