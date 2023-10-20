package qimen

// PresalesPackageConsignResponse 结构体
type PresalesPackageConsignResponse struct {
	// 响应结果:success|failure,必填
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}