package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeordersuccesscreateAPIResponse 五道口终态订单创建 API返回值
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
type AlibabawdktradeordersuccesscreateAPIResponse struct {
	model.CommonResponse
	AlibabawdktradeordersuccesscreateAPIResponseModel
}

// AlibabawdktradeordersuccesscreateAPIResponseModel is 五道口终态订单创建 成功返回结果
type AlibabawdktradeordersuccesscreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_success_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单返回结果
	OrderResult *OrderQueryResult `json:"order_result,omitempty" xml:"order_result,omitempty"`
}
