package flightuppc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightexternalalipayverifyAPIResponse 支付宝小程序验签 API返回值
// alitrip.flight.external.alipay.verify
//
// 支付宝小程序验签
type AlitripflightexternalalipayverifyAPIResponse struct {
	model.CommonResponse
	AlitripflightexternalalipayverifyAPIResponseModel
}

// AlitripflightexternalalipayverifyAPIResponseModel is 支付宝小程序验签 成功返回结果
type AlitripflightexternalalipayverifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_flight_external_alipay_verify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求失败描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否验签成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否请求成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}