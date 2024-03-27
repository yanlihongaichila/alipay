package logic

import (
	"context"
	"fmt"
	"goods/goods"
	"goods/internal/svc"
	"goods/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsStocksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsStocksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsStocksLogic {
	return &UpdateGoodsStocksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsStocksLogic) UpdateGoodsStocks(in *goods.UpdateGoodsStocksRequest) (*goods.UpdateGoodsStocksResponse, error) {
	// todo: add your logic here and delete this line
	req := map[int64]int64{}

	for _, val := range in.GoodsInfos {
		req[val.ID] = val.Num
	}

	fmt.Println("77777777777777777")
	fmt.Println(req)
	err := pkg.UpdateGoodsStocks(req)
	if err != nil {
		return nil, err
	}
	return &goods.UpdateGoodsStocksResponse{}, nil
}
