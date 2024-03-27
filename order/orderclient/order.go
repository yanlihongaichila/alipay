// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package orderclient

import (
	"context"

	"order/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateOrderGoodsRequest        = order.CreateOrderGoodsRequest
	CreateOrderGoodsResponse       = order.CreateOrderGoodsResponse
	CreateOrderRequest             = order.CreateOrderRequest
	CreateOrderResponse            = order.CreateOrderResponse
	GetOrderGoodsByOrderIDRequest  = order.GetOrderGoodsByOrderIDRequest
	GetOrderGoodsByOrderIDResponse = order.GetOrderGoodsByOrderIDResponse
	GetOrderRequest                = order.GetOrderRequest
	GetOrderResponse               = order.GetOrderResponse
	OrderGoodsInfo                 = order.OrderGoodsInfo
	OrderInfo                      = order.OrderInfo
	Request                        = order.Request
	Response                       = order.Response
	UpdateOrderRequest             = order.UpdateOrderRequest
	UpdateOrderResponse            = order.UpdateOrderResponse

	Order interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
		UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
		CreateOrderGoods(ctx context.Context, in *CreateOrderGoodsRequest, opts ...grpc.CallOption) (*CreateOrderGoodsResponse, error)
		GetOrderGoodsByOrderID(ctx context.Context, in *GetOrderGoodsByOrderIDRequest, opts ...grpc.CallOption) (*GetOrderGoodsByOrderIDResponse, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultOrder) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetOrder(ctx, in, opts...)
}

func (m *defaultOrder) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.UpdateOrder(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrderGoods(ctx context.Context, in *CreateOrderGoodsRequest, opts ...grpc.CallOption) (*CreateOrderGoodsResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrderGoods(ctx, in, opts...)
}

func (m *defaultOrder) GetOrderGoodsByOrderID(ctx context.Context, in *GetOrderGoodsByOrderIDRequest, opts ...grpc.CallOption) (*GetOrderGoodsByOrderIDResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetOrderGoodsByOrderID(ctx, in, opts...)
}
