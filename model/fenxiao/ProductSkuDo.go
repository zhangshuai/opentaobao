package fenxiao

import (
	"sync"
)

// ProductSkuDo 结构体
type ProductSkuDo struct {
	// sku的销售属性组合字符串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 关联的前端宝贝skuid
	AuctionSkuId int64 `json:"auction_sku_id,omitempty" xml:"auction_sku_id,omitempty"`
	// 代销采购价:单位分
	CostPriceFen int64 `json:"cost_price_fen,omitempty" xml:"cost_price_fen,omitempty"`
	// 经销采购价:单位分
	PriceCostDealerFen int64 `json:"price_cost_dealer_fen,omitempty" xml:"price_cost_dealer_fen,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 配额可用库存
	QuotaQuantity int64 `json:"quota_quantity,omitempty" xml:"quota_quantity,omitempty"`
	// 预扣库存
	ReservedQuantity int64 `json:"reserved_quantity,omitempty" xml:"reserved_quantity,omitempty"`
	// 关联的后端商品id
	ScitemId int64 `json:"scitem_id,omitempty" xml:"scitem_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 市场价单位分
	StandardPriceFen int64 `json:"standard_price_fen,omitempty" xml:"standard_price_fen,omitempty"`
}

var poolProductSkuDo = sync.Pool{
	New: func() any {
		return new(ProductSkuDo)
	},
}

// GetProductSkuDo() 从对象池中获取ProductSkuDo
func GetProductSkuDo() *ProductSkuDo {
	return poolProductSkuDo.Get().(*ProductSkuDo)
}

// ReleaseProductSkuDo 释放ProductSkuDo
func ReleaseProductSkuDo(v *ProductSkuDo) {
	v.Properties = ""
	v.OuterId = ""
	v.AuctionSkuId = 0
	v.CostPriceFen = 0
	v.PriceCostDealerFen = 0
	v.Quantity = 0
	v.QuotaQuantity = 0
	v.ReservedQuantity = 0
	v.ScitemId = 0
	v.SkuId = 0
	v.StandardPriceFen = 0
	poolProductSkuDo.Put(v)
}
