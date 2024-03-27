package logic

import (
	"context"
	"fmt"
	"gateWay/internal/consts"
	"gateWay/internal/pkg"
	"gateWay/internal/svc"
	"gateWay/internal/types"
	"github.com/shopspring/decimal"
	"goods/goods"
	"goods/goodsclient"
	"order/order"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	//先查询商品,判断商品是否存在,并判断库存是否够
	//传入商品的map
	reqGoodsMap := make(map[int64]int64)
	for _, val := range req.Goods {
		reqGoodsMap[val.GoodID] = val.Num
	}
	ids := []int64{}
	for _, val := range req.Goods {
		ids = append(ids, val.GoodID)
	}
	reqInfos, err := l.svcCtx.GoodsCon.GetGoodsByIDs(l.ctx, &goods.GetGoodsByIDsRequest{IDs: ids})
	if err != nil {
		return nil, err
	}
	//得到仓库中库存的map
	shopStockMap := make(map[int64]int64)
	for _, val := range reqInfos.Infos {
		shopStockMap[val.ID] = val.Stock
	}
	fmt.Println("****************")
	fmt.Println(shopStockMap)
	for key, val := range shopStockMap {
		if val-reqGoodsMap[key] < 0 {
			return nil, fmt.Errorf("%v商品库存不足", key)
		}
		reqGoodsMap[key] = -reqGoodsMap[key]
	}
	_, err = l.updatedStock(reqGoodsMap)
	if err != nil {
		return nil, err
	}
	//获取商品总价
	amount, err := l.getGoodAmount(req)
	if err != nil {
		return nil, err
	}
	orderInfo := order.OrderInfo{
		UserID:    req.UserID,
		Amount:    amount.String(),
		State:     consts.PAYMENT_WAIT,
		CreatedAt: time.Now().Unix(),
	}
	createdOrder, err := l.svcCtx.OrderCon.CreateOrder(l.ctx, &order.CreateOrderRequest{Info: &orderInfo})
	if err != nil {
		return nil, err
	}
	//添加订单商品表
	orderGoodsInfos := []*order.OrderGoodsInfo{}

	for _, val := range reqInfos.Infos {
		info := order.OrderGoodsInfo{
			OrderID:   createdOrder.Info.ID,
			GoodsID:   val.ID,
			UnitPrice: val.Price,
			GoodName:  val.Name,
			Num:       -reqGoodsMap[val.ID],
		}

		orderGoodsInfos = append(orderGoodsInfos, &info)
	}
	_, err = l.svcCtx.OrderCon.CreateOrderGoods(l.ctx, &order.CreateOrderGoodsRequest{Infos: orderGoodsInfos})
	if err != nil {
		return nil, err
	}

	url, err := pkg.GetWebPayUrl(l.svcCtx, "商品", createdOrder.Info.OrderNO, createdOrder.Info.Amount)
	if err != nil {
		return nil, err
	}

	return &types.CreateOrderResponse{Url: url}, nil
}

// 得总价
func (l *OrderCreateLogic) getGoodAmount(req *types.CreateOrderRequest) (decimal.Decimal, error) {
	var goodIDs []int64
	for _, val := range req.Goods {
		goodIDs = append(goodIDs, val.GoodID)
	}

	reqInfos, err := l.svcCtx.GoodsCon.GetGoodsByIDs(l.ctx, &goods.GetGoodsByIDsRequest{IDs: goodIDs})
	if err != nil {
		return decimal.Decimal{}, err
	}

	amount := decimal.NewFromInt(0)

	goodMap := make(map[int64]*goods.GoodsInfo)
	for _, val := range reqInfos.Infos {
		goodMap[val.ID] = val
	}

	for _, val := range req.Goods {
		goodInfo, ok := goodMap[val.GoodID]
		if !ok {
			return decimal.Decimal{}, fmt.Errorf("商品不存在")
		}

		unitPrice, err := decimal.NewFromString(goodInfo.Price)
		if err != nil {
			return decimal.Decimal{}, err
		}
		amount = amount.Add(unitPrice.Mul(decimal.NewFromInt(val.Num)))
	}

	return amount, nil
}

// 修改库存
func (l *OrderCreateLogic) updatedStock(upInfos map[int64]int64) (*goodsclient.UpdateGoodsStocksResponse, error) {
	updateGoodsInfos := goods.UpdateGoodsStocksRequest{}
	updatedReq := updateGoodsInfos.GoodsInfos

	for key, val := range upInfos {
		conInfo := goods.UpdateStockReq{
			ID:  key,
			Num: val,
		}

		updatedReq = append(updatedReq, &conInfo)
	}

	fmt.Println("99999999999999")
	fmt.Println(updatedReq)
	stocks, err := l.svcCtx.GoodsCon.UpdateGoodsStocks(l.ctx, &goods.UpdateGoodsStocksRequest{GoodsInfos: updatedReq})
	if err != nil {
		return nil, err
	}

	return stocks, nil
}
