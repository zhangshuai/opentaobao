package jym

// AuditRefundOrderDto 结构体
type AuditRefundOrderDto struct {
	// 退款单ID
	RefundOrderId string `json:"refund_order_id,omitempty" xml:"refund_order_id,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 审核结果，1：通过，0：不通过
	AuditPass int64 `json:"audit_pass,omitempty" xml:"audit_pass,omitempty"`
}