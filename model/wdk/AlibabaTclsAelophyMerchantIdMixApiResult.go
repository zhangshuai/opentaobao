package wdk

// AlibabatclsaelophymerchantidmixApiResult 结构体
type AlibabatclsaelophymerchantidmixApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回混淆id
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 获取mixId成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
