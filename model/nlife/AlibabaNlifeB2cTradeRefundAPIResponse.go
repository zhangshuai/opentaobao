package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradeRefundAPIResponse 零售+请求退款 API返回值
// alibaba.nlife.b2c.trade.refund
//
// 零售+平台请求退款接口，在零售+平台不会有资金流变动，只是订单状态的更新
type AlibabaNlifeB2cTradeRefundAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradeRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradeRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cTradeRefundAPIResponseModel).Reset()
}

// AlibabaNlifeB2cTradeRefundAPIResponseModel is 零售+请求退款 成功返回结果
type AlibabaNlifeB2cTradeRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款时间
	GmtRefund string `json:"gmt_refund,omitempty" xml:"gmt_refund,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradeRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GmtRefund = ""
	m.ExtendParams = ""
}

var poolAlibabaNlifeB2cTradeRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cTradeRefundAPIResponse)
	},
}

// GetAlibabaNlifeB2cTradeRefundAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cTradeRefundAPIResponse
func GetAlibabaNlifeB2cTradeRefundAPIResponse() *AlibabaNlifeB2cTradeRefundAPIResponse {
	return poolAlibabaNlifeB2cTradeRefundAPIResponse.Get().(*AlibabaNlifeB2cTradeRefundAPIResponse)
}

// ReleaseAlibabaNlifeB2cTradeRefundAPIResponse 将 AlibabaNlifeB2cTradeRefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cTradeRefundAPIResponse(v *AlibabaNlifeB2cTradeRefundAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cTradeRefundAPIResponse.Put(v)
}
