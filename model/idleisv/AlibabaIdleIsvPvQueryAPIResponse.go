package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvpvqueryAPIResponse 查询pv属性 API返回值
// alibaba.idle.isv.pv.query
//
// 查询pv属性
type AlibabaidleisvpvqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvpvqueryAPIResponseModel
}

// AlibabaidleisvpvqueryAPIResponseModel is 查询pv属性 成功返回结果
type AlibabaidleisvpvqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_pv_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
