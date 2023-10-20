package alitripmerchant

// AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackResponse 结构体
type AlitripmerchantgalaxyderbyvouchercardunlimitedchangecallbackResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 回调处理结果
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}