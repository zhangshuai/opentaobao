package wdk

import (
	"sync"
)

// ItemPoolActivitySkuQueryRequest 结构体
type ItemPoolActivitySkuQueryRequest struct {
	// 商品编码列表
	SkuCodes []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
	// 商品条码列表
	BarCodes []string `json:"bar_codes,omitempty" xml:"bar_codes>string,omitempty"`
	// 活动ID
	ActId string `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// erp外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 分页查询参数
	PageInfo *ActivitySkuQueryDto `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 换购品标识
	ExchangeSku bool `json:"exchange_sku,omitempty" xml:"exchange_sku,omitempty"`
}

var poolItemPoolActivitySkuQueryRequest = sync.Pool{
	New: func() any {
		return new(ItemPoolActivitySkuQueryRequest)
	},
}

// GetItemPoolActivitySkuQueryRequest() 从对象池中获取ItemPoolActivitySkuQueryRequest
func GetItemPoolActivitySkuQueryRequest() *ItemPoolActivitySkuQueryRequest {
	return poolItemPoolActivitySkuQueryRequest.Get().(*ItemPoolActivitySkuQueryRequest)
}

// ReleaseItemPoolActivitySkuQueryRequest 释放ItemPoolActivitySkuQueryRequest
func ReleaseItemPoolActivitySkuQueryRequest(v *ItemPoolActivitySkuQueryRequest) {
	v.SkuCodes = v.SkuCodes[:0]
	v.BarCodes = v.BarCodes[:0]
	v.ActId = ""
	v.OutActId = ""
	v.PageInfo = nil
	v.ExchangeSku = false
	poolItemPoolActivitySkuQueryRequest.Put(v)
}
