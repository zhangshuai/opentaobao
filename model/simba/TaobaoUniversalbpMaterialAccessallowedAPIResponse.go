package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpMaterialAccessallowedAPIResponse 物料准入判断 API返回值
// taobao.universalbp.material.accessallowed
//
// 入参推广信息，出参相关主体是否可投放。如果投放了风控不准入的商品，无法正常投放。
type TaobaoUniversalbpMaterialAccessallowedAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpMaterialAccessallowedAPIResponseModel
}

// TaobaoUniversalbpMaterialAccessallowedAPIResponseModel is 物料准入判断 成功返回结果
type TaobaoUniversalbpMaterialAccessallowedAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_material_accessallowed_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpMaterialAccessallowedTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
