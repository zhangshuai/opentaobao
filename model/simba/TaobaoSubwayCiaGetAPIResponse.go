package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayciagetAPIResponse 查询单元智能出价信息 API返回值
// taobao.subway.cia.get
//
// 查询单元智能出价信息
type TaobaosubwayciagetAPIResponse struct {
	model.CommonResponse
	TaobaosubwayciagetAPIResponseModel
}

// TaobaosubwayciagetAPIResponseModel is 查询单元智能出价信息 成功返回结果
type TaobaosubwayciagetAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_cia_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单元智能出价信息
	Result *CiaConfig `json:"result,omitempty" xml:"result,omitempty"`
}
