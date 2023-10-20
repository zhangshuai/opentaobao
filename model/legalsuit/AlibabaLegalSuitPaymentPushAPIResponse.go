package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitPaymentPushAPIResponse 外部推送缴费 API返回值
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
type AlibabaLegalSuitPaymentPushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitPaymentPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitPaymentPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitPaymentPushAPIResponseModel).Reset()
}

// AlibabaLegalSuitPaymentPushAPIResponseModel is 外部推送缴费 成功返回结果
type AlibabaLegalSuitPaymentPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_payment_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数校验错误
	ApiErrorCode string `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 参数校验错误
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 任务id
	Content int64 `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitPaymentPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiErrorCode = ""
	m.ErrorMsg = ""
	m.Content = 0
	m.ApiSuccess = false
}

var poolAlibabaLegalSuitPaymentPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitPaymentPushAPIResponse)
	},
}

// GetAlibabaLegalSuitPaymentPushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitPaymentPushAPIResponse
func GetAlibabaLegalSuitPaymentPushAPIResponse() *AlibabaLegalSuitPaymentPushAPIResponse {
	return poolAlibabaLegalSuitPaymentPushAPIResponse.Get().(*AlibabaLegalSuitPaymentPushAPIResponse)
}

// ReleaseAlibabaLegalSuitPaymentPushAPIResponse 将 AlibabaLegalSuitPaymentPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitPaymentPushAPIResponse(v *AlibabaLegalSuitPaymentPushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitPaymentPushAPIResponse.Put(v)
}
