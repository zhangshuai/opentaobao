package icbuseller

import (
	"sync"
)

// QueryTradeDto 结构体
type QueryTradeDto struct {
	// 服务code列表
	ServiceCode []string `json:"service_code,omitempty" xml:"service_code>string,omitempty"`
	// 交易id号列表
	TradeIds []int64 `json:"trade_ids,omitempty" xml:"trade_ids>int64,omitempty"`
	// 订单号列表
	OrderNos []string `json:"order_nos,omitempty" xml:"order_nos>string,omitempty"`
	// 状态列表
	Status []Null `json:"status,omitempty" xml:"status>null,omitempty"`
	// 交易单状态变更起始时间
	FireTimeStart string `json:"fire_time_start,omitempty" xml:"fire_time_start,omitempty"`
	// 交易单创建结束时间
	CreateTimeEnd string `json:"create_time_end,omitempty" xml:"create_time_end,omitempty"`
	// 交易单创建起始时间
	CreateTimeStart string `json:"create_time_start,omitempty" xml:"create_time_start,omitempty"`
	// 交易单状态变更结束时间
	FireTimeEnd string `json:"fire_time_end,omitempty" xml:"fire_time_end,omitempty"`
	// 起始值
	OffSet int64 `json:"off_set,omitempty" xml:"off_set,omitempty"`
	// 结束值
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 买家aliid
	BuyerAliId int64 `json:"buyer_ali_id,omitempty" xml:"buyer_ali_id,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否展示
	IsDisplay bool `json:"is_display,omitempty" xml:"is_display,omitempty"`
}

var poolQueryTradeDto = sync.Pool{
	New: func() any {
		return new(QueryTradeDto)
	},
}

// GetQueryTradeDto() 从对象池中获取QueryTradeDto
func GetQueryTradeDto() *QueryTradeDto {
	return poolQueryTradeDto.Get().(*QueryTradeDto)
}

// ReleaseQueryTradeDto 释放QueryTradeDto
func ReleaseQueryTradeDto(v *QueryTradeDto) {
	v.ServiceCode = v.ServiceCode[:0]
	v.TradeIds = v.TradeIds[:0]
	v.OrderNos = v.OrderNos[:0]
	v.Status = v.Status[:0]
	v.FireTimeStart = ""
	v.CreateTimeEnd = ""
	v.CreateTimeStart = ""
	v.FireTimeEnd = ""
	v.OffSet = 0
	v.Length = 0
	v.BuyerAliId = 0
	v.Page = 0
	v.PageSize = 0
	v.IsDisplay = false
	poolQueryTradeDto.Put(v)
}
