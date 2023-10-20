package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 API返回值
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.cloud.watcher.status.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponseModel).Reset()
}

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponseModel is 天猫精灵 IoT 视频 SaaS 服务-江苏电信-云回看开通状态更新 成功返回结果
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliyun_aicloud_iot_vision_saas_ctcc_jiangsu_cloud_watcher_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息。success 表示成功
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// 一次请求的唯一标识符，和请求中的 seq_id 对齐
	SeqId string `json:"seq_id,omitempty" xml:"seq_id,omitempty"`
	// 错误码。200 表示成功
	RspCode int64 `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RspMsg = ""
	m.SeqId = ""
	m.RspCode = 0
}

var poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse)
	},
}

// GetAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse
func GetAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse {
	return poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse.Get().(*AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse)
}

// ReleaseAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse 将 AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse(v *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuCloudWatcherStatusUpdateAPIResponse.Put(v)
}
