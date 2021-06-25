package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫逆向纠纷查询 APIResponse
tmall.dispute.receive.get

展示商家所有退款信息
*/
type TmallDisputeReceiveGetAPIResponse struct {
    model.CommonResponse
    Response *TmallDisputeReceiveGetResponse `json:"tmall_dispute_receive_get_response,omitempty"`
}

type TmallDisputeReceiveGetResponse struct {

    // result
    Result   *TmallDisputeReceiveGetResultSet `json:"result,omitempty"`

}