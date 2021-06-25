package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划 APIResponse
taobao.simba.campaign.update

更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
*/
type TaobaoSimbaCampaignUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignUpdateResponse `json:"taobao_simba_campaign_update_response,omitempty"`
}

type TaobaoSimbaCampaignUpdateResponse struct {

    // 修改后的推广计划
    Campaign   *Campaign `json:"campaign,omitempty"`

}