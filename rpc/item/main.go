package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"mykitex/kitex_gen/example/shop/item/itemservice"

	"log"
)

func main() {
	sc := []constant.ServerConfig{
		// 设置ip地址和端口号
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		//  设置命名空间id
		NamespaceId: "public",
		// 设置超时时间，单位毫秒
		TimeoutMs: 5000,
		// 启动时不在 CacheDir 中加载持久性 nacos 服务信息
		NotLoadCacheAtStart: true,
		//  设置日志存储文件夹
		LogDir: "/tmp/nacos/log",
		// 设置缓存存储文件夹
		CacheDir: "/tmp/nacos/cache",
		// 设置日志等级
		LogLevel: "info",
		// 设置用户名
		Username: "your-name",
		// 设置密码
		Password: "your-password",
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

	svr := itemservice.NewServer(new(ItemServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "echo"}),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
