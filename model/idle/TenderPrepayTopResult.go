package idle

// TenderPrepayTopResult 结构体
type TenderPrepayTopResult struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付宝流水号
	PrePayNo string `json:"pre_pay_no,omitempty" xml:"pre_pay_no,omitempty"`
	// 1-未支付，2-已支付
	PrePayStatus string `json:"pre_pay_status,omitempty" xml:"pre_pay_status,omitempty"`
	// 服务商appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 预付款
	PrePayAmount int64 `json:"pre_pay_amount,omitempty" xml:"pre_pay_amount,omitempty"`
	// 2-支付预付款
	PrePayAction int64 `json:"pre_pay_action,omitempty" xml:"pre_pay_action,omitempty"`
}
