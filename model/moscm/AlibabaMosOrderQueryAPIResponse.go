package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosorderqueryAPIResponse 批量查询订单信息 API返回值
// alibaba.mos.order.query
//
// 查询多笔交易信息
type AlibabamosorderqueryAPIResponse struct {
	model.CommonResponse
	AlibabamosorderqueryAPIResponseModel
}

// AlibabamosorderqueryAPIResponseModel is 批量查询订单信息 成功返回结果
type AlibabamosorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamosorderqueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
