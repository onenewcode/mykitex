// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package item

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	base "mykitex/kitex_gen/example/shop/base"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Item) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Item[number], err)
}

func (x *Item) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Item) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Item) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Item) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Stock, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetItemReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetItemReq[number], err)
}

func (x *GetItemReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetItemResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 255:
		offset, err = x.fastReadField255(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetItemResp[number], err)
}

func (x *GetItemResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Item
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Item = &v
	return offset, nil
}

func (x *GetItemResp) fastReadField255(buf []byte, _type int8) (offset int, err error) {
	var v base.BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *Item) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Item) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Item) fastWriteField2(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *Item) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *Item) fastWriteField4(buf []byte) (offset int) {
	if x.Stock == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetStock())
	return offset
}

func (x *GetItemReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetItemReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetItemResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField255(buf[offset:])
	return offset
}

func (x *GetItemResp) fastWriteField1(buf []byte) (offset int) {
	if x.Item == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetItem())
	return offset
}

func (x *GetItemResp) fastWriteField255(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 255, x.GetBaseResp())
	return offset
}

func (x *Item) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *Item) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *Item) sizeField2() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *Item) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *Item) sizeField4() (n int) {
	if x.Stock == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetStock())
	return n
}

func (x *GetItemReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetItemReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *GetItemResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField255()
	return n
}

func (x *GetItemResp) sizeField1() (n int) {
	if x.Item == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetItem())
	return n
}

func (x *GetItemResp) sizeField255() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(255, x.GetBaseResp())
	return n
}

var fieldIDToName_Item = map[int32]string{
	1: "Id",
	2: "Title",
	3: "Description",
	4: "Stock",
}

var fieldIDToName_GetItemReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_GetItemResp = map[int32]string{
	1:   "Item",
	255: "BaseResp",
}

var _ = base.File_idl_base_proto
