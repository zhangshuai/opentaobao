package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobanamadpcitemaddAPIResponse 新发商品 API返回值
// taobao.banamadpc.item.add
//
// 巴拿马供应商通过此接口新发商品
type TaobaobanamadpcitemaddAPIResponse struct {
	model.CommonResponse
	TaobaobanamadpcitemaddAPIResponseModel
}

// TaobaobanamadpcitemaddAPIResponseModel is 新发商品 成功返回结果
type TaobaobanamadpcitemaddAPIResponseModel struct {
	XMLName xml.Name `xml:"banamadpc_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	ApiResult *TaobaobanamadpcitemaddApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
