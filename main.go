package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-zookeeper/resolver"
	"log"
	"mykitex/kitex_gen/example/shop/item"
	"mykitex/kitex_gen/example/shop/item/itemservice"
	"time"
)

func main() {
	r, err := resolver.NewZookeeperResolver([]string{"121.37.143.160:2181"}, 40*time.Second)
	if err != nil {
		panic(err)
	}
	my_client, err := itemservice.NewClient("item", client.WithResolver(r))

	if err != nil {
		log.Fatal(err)
	} else {
		v, _ := my_client.GetItem(context.Background(), &item.GetItemReq{Id: 1})
		log.Println(v)
	}

}
