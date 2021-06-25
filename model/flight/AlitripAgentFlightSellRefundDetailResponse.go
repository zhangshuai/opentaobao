package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销售退票单详情 APIResponse
alitrip.agent.flight.sell.refund.detail

销售退票单详情
*/
type AlitripAgentFlightSellRefundDetailAPIResponse struct {
    model.CommonResponse
    Response *AlitripAgentFlightSellRefundDetailResponse `json:"alitrip_agent_flight_sell_refund_detail_response,omitempty"`
}

type AlitripAgentFlightSellRefundDetailResponse struct {

    // 异步获取历史数据接口返回结果
    Result   *AlitripAgentFlightSellRefundDetailResultDto `json:"result,omitempty"`

}