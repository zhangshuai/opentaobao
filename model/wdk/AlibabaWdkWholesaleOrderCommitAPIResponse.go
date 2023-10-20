package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkwholesaleordercommitAPIResponse 盒马帮采购确认订单接口 API返回值
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
type AlibabawdkwholesaleordercommitAPIResponse struct {
	model.CommonResponse
	AlibabawdkwholesaleordercommitAPIResponseModel
}

// AlibabawdkwholesaleordercommitAPIResponseModel is 盒马帮采购确认订单接口 成功返回结果
type AlibabawdkwholesaleordercommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_wholesale_order_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabawdkwholesaleordercommitApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
