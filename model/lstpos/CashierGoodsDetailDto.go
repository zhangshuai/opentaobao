package lstpos

import (
	"sync"
)

// CashierGoodsDetailDto 结构体
type CashierGoodsDetailDto struct {
	// 商品条码，可能有无码商品
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 商品数量
	Count string `json:"count,omitempty" xml:"count,omitempty"`
	// 商品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// ISV商品Id
	IsvGoodsId string `json:"isv_goods_id,omitempty" xml:"isv_goods_id,omitempty"`
	// 商品进货价 单位:分
	GoodsCostPrice int64 `json:"goods_cost_price,omitempty" xml:"goods_cost_price,omitempty"`
	// 商品销售价(折后)，单位:分
	GoodsActualPrice int64 `json:"goods_actual_price,omitempty" xml:"goods_actual_price,omitempty"`
	// 商品原价(折前)，单位:分
	GoodsOriginPrice int64 `json:"goods_origin_price,omitempty" xml:"goods_origin_price,omitempty"`
}

var poolCashierGoodsDetailDto = sync.Pool{
	New: func() any {
		return new(CashierGoodsDetailDto)
	},
}

// GetCashierGoodsDetailDto() 从对象池中获取CashierGoodsDetailDto
func GetCashierGoodsDetailDto() *CashierGoodsDetailDto {
	return poolCashierGoodsDetailDto.Get().(*CashierGoodsDetailDto)
}

// ReleaseCashierGoodsDetailDto 释放CashierGoodsDetailDto
func ReleaseCashierGoodsDetailDto(v *CashierGoodsDetailDto) {
	v.BarCode = ""
	v.Count = ""
	v.GoodsName = ""
	v.IsvGoodsId = ""
	v.GoodsCostPrice = 0
	v.GoodsActualPrice = 0
	v.GoodsOriginPrice = 0
	poolCashierGoodsDetailDto.Put(v)
}
