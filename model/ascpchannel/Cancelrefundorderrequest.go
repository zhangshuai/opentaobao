package ascpchannel

// Cancelrefundorderrequest 结构体
type Cancelrefundorderrequest struct {
	// 退款单号
	RefundNo string `json:"refund_no,omitempty" xml:"refund_no,omitempty"`
	// 外部退款单号
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
}