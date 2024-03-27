package pkg

import (
	"fmt"
	"gorm.io/gorm"
	"pkg/mysql"
)

type Goods struct {
	gorm.Model
	Name  string `gorm:"INDEX,type:varchar(20)"`
	Price string `gorm:"INDEX,type:decimal(10,2)"`
	Stock int64  `gorm:"INDEX,type:int"`
}

func NewGoods() *Goods {
	return new(Goods)
}

func GetGoodsByIDs(ids []int64) ([]Goods, error) {

	goods := NewGoods()
	goodsInfos := []Goods{}
	err := mysql.Db.Model(goods).Where("deleted_at IS NULL").Where("id IN ?", ids).Find(&goodsInfos).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return goodsInfos, nil
}

func UpdateGoodsStocks(updateGoods map[int64]int64) error {
	fmt.Println(updateGoods)
	goods := NewGoods()
	for key, val := range updateGoods {
		err := mysql.Db.Model(goods).Where("id = ?", key).Update("stock", gorm.Expr("stock  + ?", val)).Error
		if err != nil {
			return fmt.Errorf("%v商品库存修改失败", key)
		}
	}

	return nil
}
