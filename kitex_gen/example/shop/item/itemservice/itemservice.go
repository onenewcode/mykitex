// Code generated by Kitex v0.8.0. DO NOT EDIT.

package itemservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	item "mykitex/kitex_gen/example/shop/item"
)

func serviceInfo() *kitex.ServiceInfo {
	return itemServiceServiceInfo
}

var itemServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ItemService"
	handlerType := (*item.ItemService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetItem": kitex.NewMethodInfo(getItemHandler, newGetItemArgs, newGetItemResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "item",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func getItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(item.GetItemReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(item.ItemService).GetItem(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetItemArgs:
		success, err := handler.(item.ItemService).GetItem(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetItemResult)
		realResult.Success = success
	}
	return nil
}
func newGetItemArgs() interface{} {
	return &GetItemArgs{}
}

func newGetItemResult() interface{} {
	return &GetItemResult{}
}

type GetItemArgs struct {
	Req *item.GetItemReq
}

func (p *GetItemArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(item.GetItemReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetItemArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetItemArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetItemArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetItemArgs) Unmarshal(in []byte) error {
	msg := new(item.GetItemReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetItemArgs_Req_DEFAULT *item.GetItemReq

func (p *GetItemArgs) GetReq() *item.GetItemReq {
	if !p.IsSetReq() {
		return GetItemArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetItemArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetItemArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetItemResult struct {
	Success *item.GetItemResp
}

var GetItemResult_Success_DEFAULT *item.GetItemResp

func (p *GetItemResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(item.GetItemResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetItemResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetItemResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetItemResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetItemResult) Unmarshal(in []byte) error {
	msg := new(item.GetItemResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetItemResult) GetSuccess() *item.GetItemResp {
	if !p.IsSetSuccess() {
		return GetItemResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetItemResult) SetSuccess(x interface{}) {
	p.Success = x.(*item.GetItemResp)
}

func (p *GetItemResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetItemResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetItem(ctx context.Context, Req *item.GetItemReq) (r *item.GetItemResp, err error) {
	var _args GetItemArgs
	_args.Req = Req
	var _result GetItemResult
	if err = p.c.Call(ctx, "GetItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
