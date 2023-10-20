package train

import (
	"sync"
)

// TapOrderVo 结构体
type TapOrderVo struct {
	// 最晚出票时间
	LastIssueTime string `json:"last_issue_time,omitempty" xml:"last_issue_time,omitempty"`
	// 主单状态
	OrderStatusName string `json:"order_status_name,omitempty" xml:"order_status_name,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 票数
	TicketNum int64 `json:"ticket_num,omitempty" xml:"ticket_num,omitempty"`
	// 支付方式（0 现金 1 电子支付）
	VipSettleMode int64 `json:"vip_settle_mode,omitempty" xml:"vip_settle_mode,omitempty"`
	// 退改订单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 是否为紧急单
	Emergency bool `json:"emergency,omitempty" xml:"emergency,omitempty"`
}

var poolTapOrderVo = sync.Pool{
	New: func() any {
		return new(TapOrderVo)
	},
}

// GetTapOrderVo() 从对象池中获取TapOrderVo
func GetTapOrderVo() *TapOrderVo {
	return poolTapOrderVo.Get().(*TapOrderVo)
}

// ReleaseTapOrderVo 释放TapOrderVo
func ReleaseTapOrderVo(v *TapOrderVo) {
	v.LastIssueTime = ""
	v.OrderStatusName = ""
	v.TtpOrderId = 0
	v.TicketNum = 0
	v.VipSettleMode = 0
	v.TpOrderId = 0
	v.Emergency = false
	poolTapOrderVo.Put(v)
}
