package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外部交易订单取消接口 APIResponse
alibaba.wdk.trade.order.cancel

通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
*/
type AlibabaWdkTradeOrderCancelAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkTradeOrderCancelResponse `json:"alibaba_wdk_trade_order_cancel_response,omitempty"`
}

type AlibabaWdkTradeOrderCancelResponse struct {

    // 执行结果
    Result   *OrderResult `json:"result,omitempty"`

}