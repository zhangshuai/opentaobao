package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradeagreepayAPIResponse openmall订单支付 API返回值
// taobao.openmall.trade.agreepay
//
// openmall订单支付
type TaobaoopenmalltradeagreepayAPIResponse struct {
	model.CommonResponse
	TaobaoopenmalltradeagreepayAPIResponseModel
}

// TaobaoopenmalltradeagreepayAPIResponseModel is openmall订单支付 成功返回结果
type TaobaoopenmalltradeagreepayAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_agreepay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
