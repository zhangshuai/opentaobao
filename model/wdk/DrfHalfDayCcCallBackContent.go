package wdk

import (
	"sync"
)

// DrfHalfDayCcCallbackContent 结构体
type DrfHalfDayCcCallbackContent struct {
	// 子单出库关联的同城令牌
	SameTownPackages []SameTownPackage `json:"same_town_packages,omitempty" xml:"same_town_packages>same_town_package,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 缺货出库存数量
	OutOfStockStockQuantity string `json:"out_of_stock_stock_quantity,omitempty" xml:"out_of_stock_stock_quantity,omitempty"`
	// 缺货出销售数量
	OutOfStockSaleQuantity string `json:"out_of_stock_sale_quantity,omitempty" xml:"out_of_stock_sale_quantity,omitempty"`
	// 实际库存拣货数量
	ActualStockQuantity string `json:"actual_stock_quantity,omitempty" xml:"actual_stock_quantity,omitempty"`
	// 实际销售拣货数量
	ActualSaleQuantity string `json:"actual_sale_quantity,omitempty" xml:"actual_sale_quantity,omitempty"`
	// 作业内容单号
	WorkUnitContentId string `json:"work_unit_content_id,omitempty" xml:"work_unit_content_id,omitempty"`
	// 作业内容扩展属性
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 是否缺货出
	IsShortage bool `json:"is_shortage,omitempty" xml:"is_shortage,omitempty"`
}

var poolDrfHalfDayCcCallbackContent = sync.Pool{
	New: func() any {
		return new(DrfHalfDayCcCallbackContent)
	},
}

// GetDrfHalfDayCcCallbackContent() 从对象池中获取DrfHalfDayCcCallbackContent
func GetDrfHalfDayCcCallbackContent() *DrfHalfDayCcCallbackContent {
	return poolDrfHalfDayCcCallbackContent.Get().(*DrfHalfDayCcCallbackContent)
}

// ReleaseDrfHalfDayCcCallbackContent 释放DrfHalfDayCcCallbackContent
func ReleaseDrfHalfDayCcCallbackContent(v *DrfHalfDayCcCallbackContent) {
	v.SameTownPackages = v.SameTownPackages[:0]
	v.SkuName = ""
	v.SkuCode = ""
	v.OutOfStockStockQuantity = ""
	v.OutOfStockSaleQuantity = ""
	v.ActualStockQuantity = ""
	v.ActualSaleQuantity = ""
	v.WorkUnitContentId = ""
	v.Attribute = ""
	v.IsShortage = false
	poolDrfHalfDayCcCallbackContent.Put(v)
}
