package aedropshiper

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeDsOrderGetAPIResponse 买家查询订单详情 API返回值
// aliexpress.trade.ds.order.get
//
// 买家查询订单详情，用于dropshipper
type AliexpressTradeDsOrderGetAPIResponse struct {
	model.CommonResponse
	AliexpressTradeDsOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressTradeDsOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressTradeDsOrderGetAPIResponseModel).Reset()
}

// AliexpressTradeDsOrderGetAPIResponseModel is 买家查询订单详情 成功返回结果
type AliexpressTradeDsOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_trade_ds_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单信息
	Result *AeopOrderInfo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressTradeDsOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressTradeDsOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressTradeDsOrderGetAPIResponse)
	},
}

// GetAliexpressTradeDsOrderGetAPIResponse 从 sync.Pool 获取 AliexpressTradeDsOrderGetAPIResponse
func GetAliexpressTradeDsOrderGetAPIResponse() *AliexpressTradeDsOrderGetAPIResponse {
	return poolAliexpressTradeDsOrderGetAPIResponse.Get().(*AliexpressTradeDsOrderGetAPIResponse)
}

// ReleaseAliexpressTradeDsOrderGetAPIResponse 将 AliexpressTradeDsOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressTradeDsOrderGetAPIResponse(v *AliexpressTradeDsOrderGetAPIResponse) {
	v.Reset()
	poolAliexpressTradeDsOrderGetAPIResponse.Put(v)
}
