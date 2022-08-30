package ascp

// DeliveryUpsertResponse 结构体
type DeliveryUpsertResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务处理结果
	Data *DataDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}