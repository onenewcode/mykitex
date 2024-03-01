package main

import (
	"context"
	item "mykitex/kitex_gen/example/shop/item"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemReq) (resp *item.GetItemResp, err error) {
	resp = &item.GetItemResp{}
	resp.Item = &item.Item{}
	resp.Item.Id = req.GetId()
	resp.Item.Title = "Kitex"
	resp.Item.Description = "Kitex is an excellent framework!"
	return
}
