package ship

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripshipproductsyncnunberAPIResponse 船票班次变更回调 API返回值
// alitrip.ship.product.syncnunber
//
// 船票班次变更回调
type AlitripshipproductsyncnunberAPIResponse struct {
	model.CommonResponse
	AlitripshipproductsyncnunberAPIResponseModel
}

// AlitripshipproductsyncnunberAPIResponseModel is 船票班次变更回调 成功返回结果
type AlitripshipproductsyncnunberAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ship_product_syncnunber_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 成功状态
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
