package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
专属下单场景商品绑定 API请求
taobao.opentrade.special.items.bind

专属下单场景商品绑定
*/
type TaobaoOpentradeSpecialItemsBindAPIRequest struct {
    model.Params
    // 绑定专属下单场景的C端小程序ID
    _miniappId   int64
    // 本次待绑定的商品ID列表
    _itemIds   []int64
}

// 初始化TaobaoOpentradeSpecialItemsBindAPIRequest对象
func NewTaobaoOpentradeSpecialItemsBindRequest() *TaobaoOpentradeSpecialItemsBindAPIRequest{
    return &TaobaoOpentradeSpecialItemsBindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.special.items.bind"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MiniappId Setter
// 绑定专属下单场景的C端小程序ID
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) SetMiniappId(_miniappId int64) error {
    r._miniappId = _miniappId
    r.Set("miniapp_id", _miniappId)
    return nil
}

// MiniappId Getter
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetMiniappId() int64 {
    return r._miniappId
}
// ItemIds Setter
// 本次待绑定的商品ID列表
func (r *TaobaoOpentradeSpecialItemsBindAPIRequest) SetItemIds(_itemIds []int64) error {
    r._itemIds = _itemIds
    r.Set("item_ids", _itemIds)
    return nil
}

// ItemIds Getter
func (r TaobaoOpentradeSpecialItemsBindAPIRequest) GetItemIds() []int64 {
    return r._itemIds
}