package jipiao

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jipiao"
)

/* 
机票代理商】回填手续费 
taobao.alitrip.seller.refund.fillfee

回填手续费
*/
func TaobaoAlitripSellerRefundFillfee(clt *core.SDKClient, req *jipiao.TaobaoAlitripSellerRefundFillfeeRequest, session string) (*jipiao.TaobaoAlitripSellerRefundFillfeeAPIResponse, error) {
    var resp jipiao.TaobaoAlitripSellerRefundFillfeeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}