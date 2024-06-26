package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSopoPushTriggerAPIResponse 猫超共享库存寄售sopo推送触发 API返回值
// alibaba.wdk.sopo.push.trigger
//
// 猫超共享库存寄售sopo触发推送给商家
type AlibabaWdkSopoPushTriggerAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSopoPushTriggerAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSopoPushTriggerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSopoPushTriggerAPIResponseModel).Reset()
}

// AlibabaWdkSopoPushTriggerAPIResponseModel is 猫超共享库存寄售sopo推送触发 成功返回结果
type AlibabaWdkSopoPushTriggerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sopo_push_trigger_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *AlibabaWdkSopoPushTriggerApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSopoPushTriggerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSopoPushTriggerAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSopoPushTriggerAPIResponse)
	},
}

// GetAlibabaWdkSopoPushTriggerAPIResponse 从 sync.Pool 获取 AlibabaWdkSopoPushTriggerAPIResponse
func GetAlibabaWdkSopoPushTriggerAPIResponse() *AlibabaWdkSopoPushTriggerAPIResponse {
	return poolAlibabaWdkSopoPushTriggerAPIResponse.Get().(*AlibabaWdkSopoPushTriggerAPIResponse)
}

// ReleaseAlibabaWdkSopoPushTriggerAPIResponse 将 AlibabaWdkSopoPushTriggerAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSopoPushTriggerAPIResponse(v *AlibabaWdkSopoPushTriggerAPIResponse) {
	v.Reset()
	poolAlibabaWdkSopoPushTriggerAPIResponse.Put(v)
}
