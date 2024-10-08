package qimen

import (
	"sync"
)

// OrderProcessReportRequest 结构体
type OrderProcessReportRequest struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 订单信息
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 订单处理信息
	Process *Process `json:"process,omitempty" xml:"process,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenOrderprocessReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolOrderProcessReportRequest = sync.Pool{
	New: func() any {
		return new(OrderProcessReportRequest)
	},
}

// GetOrderProcessReportRequest() 从对象池中获取OrderProcessReportRequest
func GetOrderProcessReportRequest() *OrderProcessReportRequest {
	return poolOrderProcessReportRequest.Get().(*OrderProcessReportRequest)
}

// ReleaseOrderProcessReportRequest 释放OrderProcessReportRequest
func ReleaseOrderProcessReportRequest(v *OrderProcessReportRequest) {
	v.Remark = ""
	v.Order = nil
	v.Process = nil
	v.ExtendProps = nil
	poolOrderProcessReportRequest.Put(v)
}
