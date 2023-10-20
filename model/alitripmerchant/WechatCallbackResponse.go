package alitripmerchant

// WechatCallbackResponse 结构体
type WechatCallbackResponse struct {
	// 错误码，SUCCESS为清算机构接收成功，其他错误码为失败。
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息，如非空，为错误原因。
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}