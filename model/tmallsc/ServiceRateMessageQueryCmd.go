package tmallsc

import (
	"sync"
)

// ServiceRateMessageQueryCmd 结构体
type ServiceRateMessageQueryCmd struct {
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 单据类型：工单0 服务单1
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 单据id
	WorkOrderId int64 `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
	// 评价单id
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
}

var poolServiceRateMessageQueryCmd = sync.Pool{
	New: func() any {
		return new(ServiceRateMessageQueryCmd)
	},
}

// GetServiceRateMessageQueryCmd() 从对象池中获取ServiceRateMessageQueryCmd
func GetServiceRateMessageQueryCmd() *ServiceRateMessageQueryCmd {
	return poolServiceRateMessageQueryCmd.Get().(*ServiceRateMessageQueryCmd)
}

// ReleaseServiceRateMessageQueryCmd 释放ServiceRateMessageQueryCmd
func ReleaseServiceRateMessageQueryCmd(v *ServiceRateMessageQueryCmd) {
	v.ServiceCode = ""
	v.OrderType = 0
	v.WorkOrderId = 0
	v.RateId = 0
	poolServiceRateMessageQueryCmd.Put(v)
}
