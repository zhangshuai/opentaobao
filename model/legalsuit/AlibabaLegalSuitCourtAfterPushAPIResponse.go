package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtAfterPushAPIResponse 更新或者新增庭后信息 API返回值
// alibaba.legal.suit.court.after.push
//
// 供外部ISV供应商 推送庭后信息给集团诉讼
type AlibabaLegalSuitCourtAfterPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCourtAfterPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtAfterPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCourtAfterPushAPIResponseModel).Reset()
}

// AlibabaLegalSuitCourtAfterPushAPIResponseModel is 更新或者新增庭后信息 成功返回结果
type AlibabaLegalSuitCourtAfterPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_court_after_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourtAfterPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLegalSuitCourtAfterPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCourtAfterPushAPIResponse)
	},
}

// GetAlibabaLegalSuitCourtAfterPushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCourtAfterPushAPIResponse
func GetAlibabaLegalSuitCourtAfterPushAPIResponse() *AlibabaLegalSuitCourtAfterPushAPIResponse {
	return poolAlibabaLegalSuitCourtAfterPushAPIResponse.Get().(*AlibabaLegalSuitCourtAfterPushAPIResponse)
}

// ReleaseAlibabaLegalSuitCourtAfterPushAPIResponse 将 AlibabaLegalSuitCourtAfterPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCourtAfterPushAPIResponse(v *AlibabaLegalSuitCourtAfterPushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCourtAfterPushAPIResponse.Put(v)
}
