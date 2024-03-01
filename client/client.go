package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"mykitex/kitex_gen/example/shop/item"
	"mykitex/kitex_gen/example/shop/item/itemservice"
	"time"
)

func main() {
	client, err := itemservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &item.GetItemReq{Id: 1}
		resp, err := client.GetItem(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
