package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse 设备音量 API返回值
// taobao.ailab.aicloud.top.device.control.volume
//
// 设备音量
type TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopDeviceControlVolumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopDeviceControlVolumeAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopDeviceControlVolumeAPIResponseModel is 设备音量 成功返回结果
type TaobaoAilabAicloudTopDeviceControlVolumeAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_control_volume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 业务请求是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 网络请求是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopDeviceControlVolumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.Model = false
	m.IsSuccess = false
}

var poolTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse
func GetTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse() *TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse {
	return poolTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse.Get().(*TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse 将 TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse(v *TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceControlVolumeAPIResponse.Put(v)
}
