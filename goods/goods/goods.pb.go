// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: goods.proto

package goods

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
	Price string `protobuf:"bytes,30,opt,name=Price,proto3" json:"Price,omitempty"`
	Stock int64  `protobuf:"varint,40,opt,name=Stock,proto3" json:"Stock,omitempty"`
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GoodsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfo) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *GoodsInfo) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GetGoodsByIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs []int64 `protobuf:"varint,10,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
}

func (x *GetGoodsByIDsRequest) Reset() {
	*x = GetGoodsByIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsByIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsByIDsRequest) ProtoMessage() {}

func (x *GetGoodsByIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsByIDsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsByIDsRequest) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{3}
}

func (x *GetGoodsByIDsRequest) GetIDs() []int64 {
	if x != nil {
		return x.IDs
	}
	return nil
}

type GetGoodsByIDsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*GoodsInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetGoodsByIDsResponse) Reset() {
	*x = GetGoodsByIDsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsByIDsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsByIDsResponse) ProtoMessage() {}

func (x *GetGoodsByIDsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsByIDsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsByIDsResponse) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GetGoodsByIDsResponse) GetInfos() []*GoodsInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Num int64 `protobuf:"varint,20,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *UpdateStockReq) Reset() {
	*x = UpdateStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockReq) ProtoMessage() {}

func (x *UpdateStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockReq.ProtoReflect.Descriptor instead.
func (*UpdateStockReq) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStockReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateStockReq) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type UpdateGoodsStocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsInfos []*UpdateStockReq `protobuf:"bytes,10,rep,name=GoodsInfos,proto3" json:"GoodsInfos,omitempty"`
}

func (x *UpdateGoodsStocksRequest) Reset() {
	*x = UpdateGoodsStocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsStocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsStocksRequest) ProtoMessage() {}

func (x *UpdateGoodsStocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsStocksRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodsStocksRequest) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGoodsStocksRequest) GetGoodsInfos() []*UpdateStockReq {
	if x != nil {
		return x.GoodsInfos
	}
	return nil
}

type UpdateGoodsStocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*GoodsInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *UpdateGoodsStocksResponse) Reset() {
	*x = UpdateGoodsStocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsStocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsStocksResponse) ProtoMessage() {}

func (x *UpdateGoodsStocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsStocksResponse.ProtoReflect.Descriptor instead.
func (*UpdateGoodsStocksResponse) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGoodsStocksResponse) GetInfos() []*GoodsInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_goods_proto protoreflect.FileDescriptor

var file_goods_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x6e, 0x67, 0x22, 0x5b, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x49, 0x44, 0x73, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x32, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x4e, 0x75, 0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22,
	0x51, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x52, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x22, 0x43, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0xd4, 0x01, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x12, 0x27, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_goods_proto_rawDescOnce sync.Once
	file_goods_proto_rawDescData = file_goods_proto_rawDesc
)

func file_goods_proto_rawDescGZIP() []byte {
	file_goods_proto_rawDescOnce.Do(func() {
		file_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_proto_rawDescData)
	})
	return file_goods_proto_rawDescData
}

var file_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_goods_proto_goTypes = []interface{}{
	(*Request)(nil),                   // 0: goods.Request
	(*Response)(nil),                  // 1: goods.Response
	(*GoodsInfo)(nil),                 // 2: goods.GoodsInfo
	(*GetGoodsByIDsRequest)(nil),      // 3: goods.GetGoodsByIDsRequest
	(*GetGoodsByIDsResponse)(nil),     // 4: goods.GetGoodsByIDsResponse
	(*UpdateStockReq)(nil),            // 5: goods.UpdateStockReq
	(*UpdateGoodsStocksRequest)(nil),  // 6: goods.UpdateGoodsStocksRequest
	(*UpdateGoodsStocksResponse)(nil), // 7: goods.UpdateGoodsStocksResponse
}
var file_goods_proto_depIdxs = []int32{
	2, // 0: goods.GetGoodsByIDsResponse.Infos:type_name -> goods.GoodsInfo
	5, // 1: goods.UpdateGoodsStocksRequest.GoodsInfos:type_name -> goods.UpdateStockReq
	2, // 2: goods.UpdateGoodsStocksResponse.Infos:type_name -> goods.GoodsInfo
	0, // 3: goods.Goods.Ping:input_type -> goods.Request
	3, // 4: goods.Goods.GetGoodsByIDs:input_type -> goods.GetGoodsByIDsRequest
	6, // 5: goods.Goods.UpdateGoodsStocks:input_type -> goods.UpdateGoodsStocksRequest
	1, // 6: goods.Goods.Ping:output_type -> goods.Response
	4, // 7: goods.Goods.GetGoodsByIDs:output_type -> goods.GetGoodsByIDsResponse
	7, // 8: goods.Goods.UpdateGoodsStocks:output_type -> goods.UpdateGoodsStocksResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_goods_proto_init() }
func file_goods_proto_init() {
	if File_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsByIDsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsByIDsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsStocksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsStocksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_proto_goTypes,
		DependencyIndexes: file_goods_proto_depIdxs,
		MessageInfos:      file_goods_proto_msgTypes,
	}.Build()
	File_goods_proto = out.File
	file_goods_proto_rawDesc = nil
	file_goods_proto_goTypes = nil
	file_goods_proto_depIdxs = nil
}
