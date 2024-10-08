package ascpchannel

import (
	"sync"
)

// Presalesorder 结构体
type Presalesorder struct {
	// 包裹信息
	Packages []Packages `json:"packages,omitempty" xml:"packages>packages,omitempty"`
	// 订单完成时间
	OrderConfirmTime string `json:"order_confirm_time,omitempty" xml:"order_confirm_time,omitempty"`
	// 菜鸟订单号
	PresalesOrderId string `json:"presales_order_id,omitempty" xml:"presales_order_id,omitempty"`
	// 出库单号
	PresalesOrderCode string `json:"presales_order_code,omitempty" xml:"presales_order_code,omitempty"`
	// 发货仓code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 扩展属性(JSON格式)
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
}

var poolPresalesorder = sync.Pool{
	New: func() any {
		return new(Presalesorder)
	},
}

// GetPresalesorder() 从对象池中获取Presalesorder
func GetPresalesorder() *Presalesorder {
	return poolPresalesorder.Get().(*Presalesorder)
}

// ReleasePresalesorder 释放Presalesorder
func ReleasePresalesorder(v *Presalesorder) {
	v.Packages = v.Packages[:0]
	v.OrderConfirmTime = ""
	v.PresalesOrderId = ""
	v.PresalesOrderCode = ""
	v.StoreCode = ""
	v.ExtendFields = ""
	poolPresalesorder.Put(v)
}
