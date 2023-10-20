package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyordercouponvalidateAPIResponse 携带券的试单接口 API返回值
// alitrip.merchant.galaxy.order.coupon.validate
//
// 试单时可以使用优惠券，返回一个原价，一个折扣价
type AlitripmerchantgalaxyordercouponvalidateAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyordercouponvalidateAPIResponseModel
}

// AlitripmerchantgalaxyordercouponvalidateAPIResponseModel is 携带券的试单接口 成功返回结果
type AlitripmerchantgalaxyordercouponvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_order_coupon_validate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 携带优惠券试单结果
	Result *AlitripmerchantgalaxyordercouponvalidateResponse `json:"result,omitempty" xml:"result,omitempty"`
}
