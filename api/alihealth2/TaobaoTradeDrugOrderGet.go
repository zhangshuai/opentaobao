package alihealth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealth2"
)

/* 
查看订单详情 
taobao.trade.drug.order.get

商家查看订单详情
*/
func TaobaoTradeDrugOrderGet(clt *core.SDKClient, req *alihealth2.TaobaoTradeDrugOrderGetRequest, session string) (*alihealth2.TaobaoTradeDrugOrderGetAPIResponse, error) {
    var resp alihealth2.TaobaoTradeDrugOrderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}