package wdkitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdkitem"
)

/* 
商品售价会员价修改 
alibaba.wdk.item.memberprice.update

商品售价会员价修改
*/
func AlibabaWdkItemMemberpriceUpdate(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMemberpriceUpdateRequest, session string) (*wdkitem.AlibabaWdkItemMemberpriceUpdateAPIResponse, error) {
    var resp wdkitem.AlibabaWdkItemMemberpriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}