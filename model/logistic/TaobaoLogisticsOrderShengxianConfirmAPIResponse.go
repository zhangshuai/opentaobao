package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsOrderShengxianConfirmAPIResponse 物流宝生鲜冷链的发货 API返回值
// taobao.logistics.order.shengxian.confirm
//
// 优鲜送，生鲜业务使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
type TaobaoLogisticsOrderShengxianConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsOrderShengxianConfirmAPIResponseModel
}

// TaobaoLogisticsOrderShengxianConfirmAPIResponseModel is 物流宝生鲜冷链的发货 成功返回结果
type TaobaoLogisticsOrderShengxianConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_order_shengxian_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 发货成功后，返回承运商的信息
	ShipFresh *ShipFresh `json:"ship_fresh,omitempty" xml:"ship_fresh,omitempty"`
}
