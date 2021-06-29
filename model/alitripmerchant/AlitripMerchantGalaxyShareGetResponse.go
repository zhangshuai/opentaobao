package alitripmerchant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-获取小程序分享文案和图片 API返回值 
alitrip.merchant.galaxy.share.get

获取 雅高微信小程序分享素材文案和图片。
*/
type AlitripMerchantGalaxyShareGetAPIResponse struct {
    model.CommonResponse
    AlitripMerchantGalaxyShareGetResponse
}

// 星河-获取小程序分享文案和图片 成功返回结果
type AlitripMerchantGalaxyShareGetResponse struct {
    XMLName xml.Name `xml:"alitrip_merchant_galaxy_share_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 默认描述
    Result   *Response `json:"result,omitempty" xml:"result,omitempty"`
}
