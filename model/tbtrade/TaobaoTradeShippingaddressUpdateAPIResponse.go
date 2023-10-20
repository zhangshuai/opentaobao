package tbtrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeShippingaddressUpdateAPIResponse 更改交易的收货地址 API返回值
// taobao.trade.shippingaddress.update
//
// 只能更新一笔交易里面的买家收货地址 &lt;br/&gt;只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 &lt;br/&gt;更新后的发货地址可以通过taobao.trade.fullinfo.get查到 &lt;br/&gt;参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaoTradeShippingaddressUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoTradeShippingaddressUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTradeShippingaddressUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTradeShippingaddressUpdateAPIResponseModel).Reset()
}

// TaobaoTradeShippingaddressUpdateAPIResponseModel is 更改交易的收货地址 成功返回结果
type TaobaoTradeShippingaddressUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_shippingaddress_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易结构
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTradeShippingaddressUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Trade = nil
}

var poolTaobaoTradeShippingaddressUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTradeShippingaddressUpdateAPIResponse)
	},
}

// GetTaobaoTradeShippingaddressUpdateAPIResponse 从 sync.Pool 获取 TaobaoTradeShippingaddressUpdateAPIResponse
func GetTaobaoTradeShippingaddressUpdateAPIResponse() *TaobaoTradeShippingaddressUpdateAPIResponse {
	return poolTaobaoTradeShippingaddressUpdateAPIResponse.Get().(*TaobaoTradeShippingaddressUpdateAPIResponse)
}

// ReleaseTaobaoTradeShippingaddressUpdateAPIResponse 将 TaobaoTradeShippingaddressUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTradeShippingaddressUpdateAPIResponse(v *TaobaoTradeShippingaddressUpdateAPIResponse) {
	v.Reset()
	poolTaobaoTradeShippingaddressUpdateAPIResponse.Put(v)
}
