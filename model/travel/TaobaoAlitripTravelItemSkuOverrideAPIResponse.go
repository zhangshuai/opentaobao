package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelitemskuoverrideAPIResponse 【API3.0】商品级别日历价格库存修改，全量覆盖 API返回值
// taobao.alitrip.travel.item.sku.override
//
// 旅行度假新商品日历价格库存信息修改接口 第三版。提供商家通过TOP API方式修改商品sku信息。
type TaobaoalitriptravelitemskuoverrideAPIResponse struct {
	model.CommonResponse
	TaobaoalitriptravelitemskuoverrideAPIResponseModel
}

// TaobaoalitriptravelitemskuoverrideAPIResponseModel is 【API3.0】商品级别日历价格库存修改，全量覆盖 成功返回结果
type TaobaoalitriptravelitemskuoverrideAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_sku_override_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品sku修改结果
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
