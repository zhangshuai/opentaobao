package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyordergetAPIResponse 翱象拉取订单接口 API返回值
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
type AlibabaaelophyordergetAPIResponse struct {
	model.CommonResponse
	AlibabaaelophyordergetAPIResponseModel
}

// AlibabaaelophyordergetAPIResponseModel is 翱象拉取订单接口 成功返回结果
type AlibabaaelophyordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应对象
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}