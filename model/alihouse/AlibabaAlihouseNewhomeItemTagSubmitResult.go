package alihouse

// AlibabaAlihouseNewhomeItemTagSubmitResult 结构体
type AlibabaAlihouseNewhomeItemTagSubmitResult struct {
	// 错误消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 扩展信息
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 返回值
	Data *ItemTagResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功请求
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
