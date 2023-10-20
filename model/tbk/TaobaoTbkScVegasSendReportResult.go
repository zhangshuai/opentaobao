package tbk

// TaobaotbkscvegassendreportResult 结构体
type TaobaotbkscvegassendreportResult struct {
	// msgInfo
	Msginfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	Msgcode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *RightsSendRptDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
