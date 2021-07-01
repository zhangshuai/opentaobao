package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品发布页 API请求
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
type TaobaoBanamadpcItemEditRenderAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TaobaoBanamadpcItemEditRenderAPIRequest对象
func NewTaobaoBanamadpcItemEditRenderRequest() *TaobaoBanamadpcItemEditRenderAPIRequest{
    return &TaobaoBanamadpcItemEditRenderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetApiMethodName() string {
    return "taobao.banamadpc.item.edit.render"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoBanamadpcItemEditRenderAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoBanamadpcItemEditRenderAPIRequest) GetItemId() int64 {
    return r._itemId
}