package main

import (
	"log"
	item "mykitex/kitex_gen/example/shop/item/itemservice"
)

func main() {
	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
