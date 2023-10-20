package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelgereralproductupdateAPIResponse 通用类目产品发布编辑 API返回值
// alitrip.travel.gereralproduct.update
//
// 提供给飞猪供销平台供应商发布编辑通用类目产品的API
type AlitriptravelgereralproductupdateAPIResponse struct {
	model.CommonResponse
	AlitriptravelgereralproductupdateAPIResponseModel
}

// AlitriptravelgereralproductupdateAPIResponseModel is 通用类目产品发布编辑 成功返回结果
type AlitriptravelgereralproductupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_gereralproduct_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// firstResult
	FirstResult *TopTravelItem `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
