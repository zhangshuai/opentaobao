package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家取消店铺vip的优惠 API请求
taobao.crm.shopvip.cancel

此接口用于取消VIP优惠
*/
type TaobaoCrmShopvipCancelRequest struct {
    model.Params
}

// 初始化TaobaoCrmShopvipCancelRequest对象
func NewTaobaoCrmShopvipCancelRequest() *TaobaoCrmShopvipCancelRequest{
    return &TaobaoCrmShopvipCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmShopvipCancelRequest) GetApiMethodName() string {
    return "taobao.crm.shopvip.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmShopvipCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
