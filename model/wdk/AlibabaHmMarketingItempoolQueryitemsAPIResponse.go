package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolqueryitemsAPIResponse 查询商品池活动下面的商品 API返回值
// alibaba.hm.marketing.itempool.queryitems
//
// 查询商品池活动下面的商品
type AlibabahmmarketingitempoolqueryitemsAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingitempoolqueryitemsAPIResponseModel
}

// AlibabahmmarketingitempoolqueryitemsAPIResponseModel is 查询商品池活动下面的商品 成功返回结果
type AlibabahmmarketingitempoolqueryitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_queryitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}