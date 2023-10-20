package axindata

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelfscrouteapiproductofflineAPIResponse 产品下线 API返回值
// taobao.alitrip.travel.fsc.route.api.product.offline
//
// 产品下线
type TaobaoalitriptravelfscrouteapiproductofflineAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelfscrouteapiproductofflineAPIResponseModel
}

// TaobaoalitriptravelfscrouteapiproductofflineAPIResponseModel is 产品下线 成功返回结果
type TaobaoalitriptravelfscrouteapiproductofflineAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_fsc_route_api_product_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通用返回结果
	TopResult *TaobaoalitriptravelfscrouteapiproductofflineTopResult `json:"top_result,omitempty" xml:"top_result,omitempty"`
}