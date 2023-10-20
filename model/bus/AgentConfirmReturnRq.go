package bus

import (
	"sync"
)

// AgentConfirmReturnRq 结构体
type AgentConfirmReturnRq struct {
	// 商家单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 按票退 时的票号
	AgentTicketId string `json:"agent_ticket_id,omitempty" xml:"agent_ticket_id,omitempty"`
	// 调用方表示
	ComeFrom string `json:"come_from,omitempty" xml:"come_from,omitempty"`
	// 按票退 时的乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 参考退款金额（以实际为准）
	RefundAmount string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 签名
	Isign string `json:"isign,omitempty" xml:"isign,omitempty"`
	// 退票成功时间点
	SuccessTime string `json:"success_time,omitempty" xml:"success_time,omitempty"`
	// 退票结果 1-成功 2-失败
	AgentReturnTicketStatus int64 `json:"agent_return_ticket_status,omitempty" xml:"agent_return_ticket_status,omitempty"`
	// 退票类型 0-按票退 1-按单退
	AgentReturnTicketType int64 `json:"agent_return_ticket_type,omitempty" xml:"agent_return_ticket_type,omitempty"`
}

var poolAgentConfirmReturnRq = sync.Pool{
	New: func() any {
		return new(AgentConfirmReturnRq)
	},
}

// GetAgentConfirmReturnRq() 从对象池中获取AgentConfirmReturnRq
func GetAgentConfirmReturnRq() *AgentConfirmReturnRq {
	return poolAgentConfirmReturnRq.Get().(*AgentConfirmReturnRq)
}

// ReleaseAgentConfirmReturnRq 释放AgentConfirmReturnRq
func ReleaseAgentConfirmReturnRq(v *AgentConfirmReturnRq) {
	v.AgentOrderId = ""
	v.AgentTicketId = ""
	v.ComeFrom = ""
	v.PassengerName = ""
	v.RefundAmount = ""
	v.Isign = ""
	v.SuccessTime = ""
	v.AgentReturnTicketStatus = 0
	v.AgentReturnTicketType = 0
	poolAgentConfirmReturnRq.Put(v)
}
