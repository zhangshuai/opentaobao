package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemqueryAPIResponse 服务商闲鱼商品查询 API返回值
// alibaba.idle.isv.item.query
//
// 服务商ISV闲鱼商品查询
type AlibabaidleisvitemqueryAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvitemqueryAPIResponseModel
}

// AlibabaidleisvitemqueryAPIResponseModel is 服务商闲鱼商品查询 成功返回结果
type AlibabaidleisvitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果result
	Result *AlibabaidleisvitemqueryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
