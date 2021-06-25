package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建退款留言/凭证 APIResponse
taobao.refund.message.add

创建退款留言/凭证
*/
type TaobaoRefundMessageAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRefundMessageAddResponse `json:"taobao_refund_message_add_response,omitempty"`
}

type TaobaoRefundMessageAddResponse struct {

    // 退款信息。包含id和created
    RefundMessage   *RefundMessage `json:"refund_message,omitempty"`

}