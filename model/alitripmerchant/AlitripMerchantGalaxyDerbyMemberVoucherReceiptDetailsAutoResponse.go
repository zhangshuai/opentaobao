package alitripmerchant

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoResponse 结构体
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoResponse struct {
	// 返回类型
	Content []DerbyVoucherReceiptAutoVo `json:"content,omitempty" xml:"content>derby_voucher_receipt_auto_vo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}