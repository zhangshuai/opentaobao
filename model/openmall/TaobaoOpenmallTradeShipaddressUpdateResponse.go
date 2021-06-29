package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Openmall订单收货地址修改 API返回值 
taobao.openmall.trade.shipaddress.update

Openmall订单收货地址修改
*/
type TaobaoOpenmallTradeShipaddressUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeShipaddressUpdateResponse
}

// Openmall订单收货地址修改 成功返回结果
type TaobaoOpenmallTradeShipaddressUpdateResponse struct {
    XMLName xml.Name `xml:"openmall_trade_shipaddress_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 订单号
    Tid   string `json:"tid,omitempty" xml:"tid,omitempty"`
}