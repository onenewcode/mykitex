package main

import (
	"context"
	"mykitex/kitex_gen/example/shop/item"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"mykitex/kitex_gen/example/shop/item/itemservice"
	"time"
)

var (
	cli itemservice.Client
)

func main() {
	c, err := itemservice.NewClient("example.shop.item", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	cli = c

	hz := server.New(server.WithHostPorts("localhost:8889"))

	hz.GET("/api/item", Handler)

	if err := hz.Run(); err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx context.Context, c *app.RequestContext) {
	req := &item.GetItemReq{Id: 1}
	req.Id = 1024
	resp, err := cli.GetItem(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	c.String(200, resp.String())
}
