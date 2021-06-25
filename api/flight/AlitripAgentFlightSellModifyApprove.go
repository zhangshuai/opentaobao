package flight

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/flight"
)

/* 
销售改签确认 
alitrip.agent.flight.sell.modify.approve

销售改签确认
*/
func AlitripAgentFlightSellModifyApprove(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyApproveRequest, session string) (*flight.AlitripAgentFlightSellModifyApproveResponse, error) {
    var resp flight.AlitripAgentFlightSellModifyApproveAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}