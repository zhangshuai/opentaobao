package ascp

// AlibabaAscpSuborderEstcontimeModifyResult 结构体
type AlibabaAscpSuborderEstcontimeModifyResult struct {
	// 对用户展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误代码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 系统内部错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 时效修改返回信息
	Value *EstConTimeModifyDto `json:"value,omitempty" xml:"value,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
