package trade

import (
	"sync"
)

// FastBuyPosItemBo 结构体
type FastBuyPosItemBo struct {
	// 购买库存数量
	InvQuantity string `json:"inv_quantity,omitempty" xml:"inv_quantity,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商家商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 扩展属性
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 商品单价
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 商品行号,单次下单自增
	LineId int64 `json:"line_id,omitempty" xml:"line_id,omitempty"`
	// 一阶段优惠金额（商品实际优惠金额）：单价*数量-商家优惠
	PromotionFee int64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
}

var poolFastBuyPosItemBo = sync.Pool{
	New: func() any {
		return new(FastBuyPosItemBo)
	},
}

// GetFastBuyPosItemBo() 从对象池中获取FastBuyPosItemBo
func GetFastBuyPosItemBo() *FastBuyPosItemBo {
	return poolFastBuyPosItemBo.Get().(*FastBuyPosItemBo)
}

// ReleaseFastBuyPosItemBo 释放FastBuyPosItemBo
func ReleaseFastBuyPosItemBo(v *FastBuyPosItemBo) {
	v.InvQuantity = ""
	v.BarCode = ""
	v.Title = ""
	v.SkuCode = ""
	v.ExtendInfo = ""
	v.UnitPrice = 0
	v.LineId = 0
	v.PromotionFee = 0
	poolFastBuyPosItemBo.Put(v)
}
