package main

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	zkregistry "github.com/kitex-contrib/registry-zookeeper/registry"
	"log"
	"mykitex/kitex_gen/example/shop/item/itemservice"
	"net"
	"time"
)

func main() {
	r, err := zkregistry.NewZookeeperRegistry([]string{"121.37.143.160:2181"}, 40*time.Second)
	if err != nil {
		panic(err)
	}

	var svr = itemservice.NewServer(new(ItemServiceImpl), server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "item", Addr: &net.TCPAddr{IP: []byte("127.0.0.1"), Port: 9999}})) //设置服务注册在zk中的信息
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
