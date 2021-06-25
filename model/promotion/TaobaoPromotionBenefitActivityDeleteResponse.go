package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除关联的活动权益 APIResponse
taobao.promotion.benefit.activity.delete

删除关联的活动权益
*/
type TaobaoPromotionBenefitActivityDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionBenefitActivityDeleteResponse `json:"taobao_promotion_benefit_activity_delete_response,omitempty"`
}

type TaobaoPromotionBenefitActivityDeleteResponse struct {

    // 删除是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}