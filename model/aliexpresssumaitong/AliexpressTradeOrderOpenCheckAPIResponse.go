package aliexpresssumaitong

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeOrderOpenCheckAPIResponse Aliexpress开放平台下单前置检查 API返回值
// aliexpress.trade.order.open.check
//
// Aliexpress开放平台下单前通过下单入参获取token
type AliexpressTradeOrderOpenCheckAPIResponse struct {
	model.CommonResponse
	AliexpressTradeOrderOpenCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressTradeOrderOpenCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressTradeOrderOpenCheckAPIResponseModel).Reset()
}

// AliexpressTradeOrderOpenCheckAPIResponseModel is Aliexpress开放平台下单前置检查 成功返回结果
type AliexpressTradeOrderOpenCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_trade_order_open_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 预下单返回值
	Result *PreCheckResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressTradeOrderOpenCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressTradeOrderOpenCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressTradeOrderOpenCheckAPIResponse)
	},
}

// GetAliexpressTradeOrderOpenCheckAPIResponse 从 sync.Pool 获取 AliexpressTradeOrderOpenCheckAPIResponse
func GetAliexpressTradeOrderOpenCheckAPIResponse() *AliexpressTradeOrderOpenCheckAPIResponse {
	return poolAliexpressTradeOrderOpenCheckAPIResponse.Get().(*AliexpressTradeOrderOpenCheckAPIResponse)
}

// ReleaseAliexpressTradeOrderOpenCheckAPIResponse 将 AliexpressTradeOrderOpenCheckAPIResponse 保存到 sync.Pool
func ReleaseAliexpressTradeOrderOpenCheckAPIResponse(v *AliexpressTradeOrderOpenCheckAPIResponse) {
	v.Reset()
	poolAliexpressTradeOrderOpenCheckAPIResponse.Put(v)
}
