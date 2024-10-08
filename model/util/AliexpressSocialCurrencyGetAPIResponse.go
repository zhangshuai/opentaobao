package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialCurrencyGetAPIResponse 币种获取接口 API返回值
// aliexpress.social.currency.get
//
// 获取目前AE社交支持的币种
type AliexpressSocialCurrencyGetAPIResponse struct {
	model.CommonResponse
	AliexpressSocialCurrencyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSocialCurrencyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialCurrencyGetAPIResponseModel).Reset()
}

// AliexpressSocialCurrencyGetAPIResponseModel is 币种获取接口 成功返回结果
type AliexpressSocialCurrencyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_currency_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包类型
	Result *ItemPickPagingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialCurrencyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSocialCurrencyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialCurrencyGetAPIResponse)
	},
}

// GetAliexpressSocialCurrencyGetAPIResponse 从 sync.Pool 获取 AliexpressSocialCurrencyGetAPIResponse
func GetAliexpressSocialCurrencyGetAPIResponse() *AliexpressSocialCurrencyGetAPIResponse {
	return poolAliexpressSocialCurrencyGetAPIResponse.Get().(*AliexpressSocialCurrencyGetAPIResponse)
}

// ReleaseAliexpressSocialCurrencyGetAPIResponse 将 AliexpressSocialCurrencyGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialCurrencyGetAPIResponse(v *AliexpressSocialCurrencyGetAPIResponse) {
	v.Reset()
	poolAliexpressSocialCurrencyGetAPIResponse.Put(v)
}
