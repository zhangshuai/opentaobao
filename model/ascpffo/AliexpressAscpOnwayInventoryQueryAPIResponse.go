package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressascponwayinventoryqueryAPIResponse AliExpress在途库存查询API API返回值
// aliexpress.ascp.onway.inventory.query
//
// AliExpress在途库存查询API
type AliexpressascponwayinventoryqueryAPIResponse struct {
	model.CommonResponse
	AliexpressascponwayinventoryqueryAPIResponseModel
}

// AliexpressascponwayinventoryqueryAPIResponseModel is AliExpress在途库存查询API 成功返回结果
type AliexpressascponwayinventoryqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ascp_onway_inventory_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressascponwayinventoryqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
