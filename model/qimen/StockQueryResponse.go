package qimen

import (
	"sync"
)

// StockQueryResponse 结构体
type StockQueryResponse struct {
	// 商品的库存信息列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数
	TotalCount int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

var poolStockQueryResponse = sync.Pool{
	New: func() any {
		return new(StockQueryResponse)
	},
}

// GetStockQueryResponse() 从对象池中获取StockQueryResponse
func GetStockQueryResponse() *StockQueryResponse {
	return poolStockQueryResponse.Get().(*StockQueryResponse)
}

// ReleaseStockQueryResponse 释放StockQueryResponse
func ReleaseStockQueryResponse(v *StockQueryResponse) {
	v.Items = v.Items[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.TotalCount = 0
	poolStockQueryResponse.Put(v)
}
