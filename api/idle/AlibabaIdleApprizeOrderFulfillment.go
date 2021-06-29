package idle

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idle"
)

/* 
鉴定担保资金订单履约 
alibaba.idle.apprize.order.fulfillment

服务商针对自己的服务订单进行履约
*/
func AlibabaIdleApprizeOrderFulfillment(clt *core.SDKClient, req *idle.AlibabaIdleApprizeOrderFulfillmentRequest, session string) (*idle.AlibabaIdleApprizeOrderFulfillmentAPIResponse, error) {
    var resp idle.AlibabaIdleApprizeOrderFulfillmentAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}