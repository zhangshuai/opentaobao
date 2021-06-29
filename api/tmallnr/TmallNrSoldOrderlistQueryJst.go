package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
已入塔商家查询订单列表 
tmall.nr.sold.orderlist.query.jst

该服务用于已入聚石塔的商家，获取订单列表信息；
*/
func TmallNrSoldOrderlistQueryJst(clt *core.SDKClient, req *tmallnr.TmallNrSoldOrderlistQueryJstRequest, session string) (*tmallnr.TmallNrSoldOrderlistQueryJstAPIResponse, error) {
    var resp tmallnr.TmallNrSoldOrderlistQueryJstAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}