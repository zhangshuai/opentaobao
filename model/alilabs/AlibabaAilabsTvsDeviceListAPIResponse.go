package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTvsDeviceListAPIResponse 获取TVS设备列表 API返回值
// alibaba.ailabs.tvs.device.list
//
// 获取用户所绑定的TVS设备列表
type AlibabaAilabsTvsDeviceListAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTvsDeviceListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTvsDeviceListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTvsDeviceListAPIResponseModel).Reset()
}

// AlibabaAilabsTvsDeviceListAPIResponseModel is 获取TVS设备列表 成功返回结果
type AlibabaAilabsTvsDeviceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tvs_device_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备信息列表
	Devices []AlibabaAilabsTvsDeviceListData `json:"devices,omitempty" xml:"devices>alibaba_ailabs_tvs_device_list_data,omitempty"`
	// 接口调用结果代码，200代表调用成功。
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 接口调用错误时给出的错误相关信息。
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 服务请求是否成功。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTvsDeviceListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Devices = m.Devices[:0]
	m.StatusCode = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolAlibabaAilabsTvsDeviceListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTvsDeviceListAPIResponse)
	},
}

// GetAlibabaAilabsTvsDeviceListAPIResponse 从 sync.Pool 获取 AlibabaAilabsTvsDeviceListAPIResponse
func GetAlibabaAilabsTvsDeviceListAPIResponse() *AlibabaAilabsTvsDeviceListAPIResponse {
	return poolAlibabaAilabsTvsDeviceListAPIResponse.Get().(*AlibabaAilabsTvsDeviceListAPIResponse)
}

// ReleaseAlibabaAilabsTvsDeviceListAPIResponse 将 AlibabaAilabsTvsDeviceListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTvsDeviceListAPIResponse(v *AlibabaAilabsTvsDeviceListAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTvsDeviceListAPIResponse.Put(v)
}
