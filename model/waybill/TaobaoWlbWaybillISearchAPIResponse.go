package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillISearchAPIResponse 查询面单服务订购及面单使用情况v1.0 API返回值
// taobao.wlb.waybill.i.search
//
// 获取发货地&CP开通状态&账户的使用情况
type TaobaoWlbWaybillISearchAPIResponse struct {
	model.CommonResponse
	TaobaoWlbWaybillISearchAPIResponseModel
}

// TaobaoWlbWaybillISearchAPIResponseModel is 查询面单服务订购及面单使用情况v1.0 成功返回结果
type TaobaoWlbWaybillISearchAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_waybill_i_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订购关系
	Subscribtions []WaybillApplySubscriptionInfo `json:"subscribtions,omitempty" xml:"subscribtions>waybill_apply_subscription_info,omitempty"`
}
