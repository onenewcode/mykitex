package main

import (
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/v2/resolver"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"mykitex/kitex_gen/example/shop/item/itemservice"
)

var (
	cli itemservice.Client
)

func main() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		Username:            "your-name",
		Password:            "your-password",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	my_client, err := itemservice.NewClient("echo", client.WithResolver(resolver.NewNacosResolver(cli)))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(my_client)

}
