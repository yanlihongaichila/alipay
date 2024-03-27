// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package goodsclient

import (
	"context"
	"goods/goods"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetGoodsByIDsRequest      = goods.GetGoodsByIDsRequest
	GetGoodsByIDsResponse     = goods.GetGoodsByIDsResponse
	GoodsInfo                 = goods.GoodsInfo
	Request                   = goods.Request
	Response                  = goods.Response
	UpdateGoodsStocksRequest  = goods.UpdateGoodsStocksRequest
	UpdateGoodsStocksResponse = goods.UpdateGoodsStocksResponse
	UpdateStockReq            = goods.UpdateStockReq

	Goods interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		GetGoodsByIDs(ctx context.Context, in *GetGoodsByIDsRequest, opts ...grpc.CallOption) (*GetGoodsByIDsResponse, error)
		UpdateGoodsStocks(ctx context.Context, in *UpdateGoodsStocksRequest, opts ...grpc.CallOption) (*UpdateGoodsStocksResponse, error)
	}

	defaultGoods struct {
		cli zrpc.Client
	}
)

func NewGoods(cli zrpc.Client) Goods {
	return &defaultGoods{
		cli: cli,
	}
}

func (m *defaultGoods) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := goods.NewGoodsClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultGoods) GetGoodsByIDs(ctx context.Context, in *GetGoodsByIDsRequest, opts ...grpc.CallOption) (*GetGoodsByIDsResponse, error) {
	client := goods.NewGoodsClient(m.cli.Conn())
	return client.GetGoodsByIDs(ctx, in, opts...)
}

func (m *defaultGoods) UpdateGoodsStocks(ctx context.Context, in *UpdateGoodsStocksRequest, opts ...grpc.CallOption) (*UpdateGoodsStocksResponse, error) {
	client := goods.NewGoodsClient(m.cli.Conn())
	return client.UpdateGoodsStocks(ctx, in, opts...)
}
