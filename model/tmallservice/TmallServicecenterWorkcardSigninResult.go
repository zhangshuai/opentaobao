package tmallservice

// TmallservicecenterworkcardsigninResult 结构体
type TmallservicecenterworkcardsigninResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
