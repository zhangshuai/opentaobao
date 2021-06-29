package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-取消预订 API返回值 
alitrip.merchant.galaxy.order.cancel

雅高酒店用户使用该接口，取消酒店预订
*/
type AlitripMerchantGalaxyOrderCancelAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyOrderCancelResponse
}

// 星河-取消预订 成功返回结果
type AlitripMerchantGalaxyOrderCancelResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 默认描述
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`
}
