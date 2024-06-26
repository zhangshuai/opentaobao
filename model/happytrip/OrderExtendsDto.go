package happytrip

import (
	"sync"
)

// OrderExtendsDto 结构体
type OrderExtendsDto struct {
	// 供应商拒绝原因
	AgentFail string `json:"agent_fail,omitempty" xml:"agent_fail,omitempty"`
	// 买家拒绝原因
	BuyerFail string `json:"buyer_fail,omitempty" xml:"buyer_fail,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorsCode string `json:"errors_code,omitempty" xml:"errors_code,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 拒绝选择最低价的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 旅行目的
	TripPurpose string `json:"trip_purpose,omitempty" xml:"trip_purpose,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 所属订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolOrderExtendsDto = sync.Pool{
	New: func() any {
		return new(OrderExtendsDto)
	},
}

// GetOrderExtendsDto() 从对象池中获取OrderExtendsDto
func GetOrderExtendsDto() *OrderExtendsDto {
	return poolOrderExtendsDto.Get().(*OrderExtendsDto)
}

// ReleaseOrderExtendsDto 释放OrderExtendsDto
func ReleaseOrderExtendsDto(v *OrderExtendsDto) {
	v.AgentFail = ""
	v.BuyerFail = ""
	v.ErrorMessage = ""
	v.ErrorsCode = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Reason = ""
	v.TripPurpose = ""
	v.Id = 0
	v.OrderId = 0
	poolOrderExtendsDto.Put(v)
}
