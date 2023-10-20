package drug

import (
	"sync"
)

// SubOrderDto 结构体
type SubOrderDto struct {
	// 套装商品子商品系信息
	SuitSubItemDtoList []SuitSubItemDto `json:"suit_sub_item_dto_list,omitempty" xml:"suit_sub_item_dto_list>suit_sub_item_dto,omitempty"`
	// 商品单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商家自定义商品ID
	OutItemdId string `json:"out_itemd_id,omitempty" xml:"out_itemd_id,omitempty"`
	// 子订单ID
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品营销类型：0-普通，1-满就送；2-兑换券
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 是否为处方药
	Rx int64 `json:"rx,omitempty" xml:"rx,omitempty"`
	// 商品的skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolSubOrderDto = sync.Pool{
	New: func() any {
		return new(SubOrderDto)
	},
}

// GetSubOrderDto() 从对象池中获取SubOrderDto
func GetSubOrderDto() *SubOrderDto {
	return poolSubOrderDto.Get().(*SubOrderDto)
}

// ReleaseSubOrderDto 释放SubOrderDto
func ReleaseSubOrderDto(v *SubOrderDto) {
	v.SuitSubItemDtoList = v.SuitSubItemDtoList[:0]
	v.Unit = ""
	v.Title = ""
	v.OutItemdId = ""
	v.SubOrderId = 0
	v.ItemId = 0
	v.Type = 0
	v.Price = 0
	v.BuyAmount = 0
	v.Rx = 0
	v.SkuId = 0
	poolSubOrderDto.Put(v)
}
