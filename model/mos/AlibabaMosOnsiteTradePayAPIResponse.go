package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradePayAPIResponse 新商场当面付商户扫码付 API返回值
// alibaba.mos.onsite.trade.pay
//
// 收银员使用扫码设备读取用户“付款码”后，将二维码或条码信息通过本接口上送至喵街发起支付。
type AlibabaMosOnsiteTradePayAPIResponse struct {
	model.CommonResponse
	AlibabaMosOnsiteTradePayAPIResponseModel
}

// AlibabaMosOnsiteTradePayAPIResponseModel is 新商场当面付商户扫码付 成功返回结果
type AlibabaMosOnsiteTradePayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_onsite_trade_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建订单响应信息。必然返回
	OnsiteTradePayResponse *OnsiteTradePayResponse `json:"onsite_trade_pay_response,omitempty" xml:"onsite_trade_pay_response,omitempty"`
}
