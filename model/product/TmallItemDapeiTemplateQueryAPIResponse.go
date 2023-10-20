package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemdapeitemplatequeryAPIResponse 搭配查询接口 API返回值
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
type TmallitemdapeitemplatequeryAPIResponse struct {
	model.CommonResponse
	TmallitemdapeitemplatequeryAPIResponseModel
}

// TmallitemdapeitemplatequeryAPIResponseModel is 搭配查询接口 成功返回结果
type TmallitemdapeitemplatequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_dapei_template_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallitemdapeitemplatequeryResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
