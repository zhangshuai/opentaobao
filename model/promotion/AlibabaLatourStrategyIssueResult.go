package promotion

// AlibabalatourstrategyissueResult 结构体
type AlibabalatourstrategyissueResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回结果
	Data *StrategyIssueResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
