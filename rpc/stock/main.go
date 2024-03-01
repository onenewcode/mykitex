package main

import (
	"log"
	stock "mykitex/kitex_gen/example/shop/stock/stockservice"
)

func main() {
	svr := stock.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
