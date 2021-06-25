package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售改签拒绝 APIResponse
alitrip.agent.flight.sell.modify.refuse

销售改签拒绝
*/
type AlitripAgentFlightSellModifyRefuseAPIResponse struct {
    model.CommonResponse
    Response *AlitripAgentFlightSellModifyRefuseResponse `json:"alitrip_agent_flight_sell_modify_refuse_response,omitempty"`
}

type AlitripAgentFlightSellModifyRefuseResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellModifyRefuseResultDto `json:"result,omitempty"`

}