package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询无条件单品优惠活动列表 APIResponse
taobao.promotionmisc.item.activity.list.get

查询无条件单品优惠活动列表
*/
type TaobaoPromotionmiscItemActivityListGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionmiscItemActivityListGetResponse `json:"taobao_promotionmisc_item_activity_list_get_response,omitempty"`
}

type TaobaoPromotionmiscItemActivityListGetResponse struct {

    // 只包含活动的主要信息，如activity_id，name，description，start_time，end_time，participate_range。优惠的其他详细信息，请通过taobao.promotionmisc.item.activity.get获取。
    ItemPromotionList   []ItemPromotion `json:"item_promotion_list,omitempty"`

    // 记录总条数。
    TotalCount   int64 `json:"total_count,omitempty"`

}