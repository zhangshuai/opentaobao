package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWholesaleGoodsSearchAPIResponse 批发市场产品搜索 API返回值
// alibaba.wholesale.goods.search
//
// 批发市场产品搜索
type AlibabaWholesaleGoodsSearchAPIResponse struct {
	model.CommonResponse
	AlibabaWholesaleGoodsSearchAPIResponseModel
}

// AlibabaWholesaleGoodsSearchAPIResponseModel is 批发市场产品搜索 成功返回结果
type AlibabaWholesaleGoodsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wholesale_goods_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 在线批发商品搜索结果
	WholesaleGoodsSearchResult *WholesaleSearchOpenResult `json:"wholesale_goods_search_result,omitempty" xml:"wholesale_goods_search_result,omitempty"`
}
