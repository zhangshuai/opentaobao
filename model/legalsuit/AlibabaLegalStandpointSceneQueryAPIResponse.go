package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalStandpointSceneQueryAPIResponse 查询场景 API返回值
// alibaba.legal.standpoint.scene.query
//
// 查询场景
type AlibabaLegalStandpointSceneQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLegalStandpointSceneQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointSceneQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalStandpointSceneQueryAPIResponseModel).Reset()
}

// AlibabaLegalStandpointSceneQueryAPIResponseModel is 查询场景 成功返回结果
type AlibabaLegalStandpointSceneQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_scene_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 场景集合
	Content []SceneOption `json:"content,omitempty" xml:"content>scene_option,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalStandpointSceneQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = m.Content[:0]
	m.ErrorMsg = ""
	m.ErrorCodeRes = 0
	m.SuccessRes = false
}

var poolAlibabaLegalStandpointSceneQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalStandpointSceneQueryAPIResponse)
	},
}

// GetAlibabaLegalStandpointSceneQueryAPIResponse 从 sync.Pool 获取 AlibabaLegalStandpointSceneQueryAPIResponse
func GetAlibabaLegalStandpointSceneQueryAPIResponse() *AlibabaLegalStandpointSceneQueryAPIResponse {
	return poolAlibabaLegalStandpointSceneQueryAPIResponse.Get().(*AlibabaLegalStandpointSceneQueryAPIResponse)
}

// ReleaseAlibabaLegalStandpointSceneQueryAPIResponse 将 AlibabaLegalStandpointSceneQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalStandpointSceneQueryAPIResponse(v *AlibabaLegalStandpointSceneQueryAPIResponse) {
	v.Reset()
	poolAlibabaLegalStandpointSceneQueryAPIResponse.Put(v)
}
