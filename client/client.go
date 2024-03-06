package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"mykitex/kitex_gen/example/shop/item"
	"mykitex/kitex_gen/example/shop/item/itemservice"
)

func main() {
	// 使用时请传入真实 etcd 的服务地址，本例中为 127.0.0.1:2379
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	// 指定 Resolve
	cl, err := itemservice.NewClient("example.shop.item",
		client.WithResolver(r),
		client.WithClientBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "example.shop.item",
			},
		),
	)
	if err != nil {
		log.Println("发现服务失败")
	}
	p, _ := cl.GetItem(context.Background(), &item.GetItemReq{Id: 1})
	log.Println(p)
}
