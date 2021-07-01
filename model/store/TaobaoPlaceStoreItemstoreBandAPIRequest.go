package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店商品关联绑定接口 API请求
taobao.place.store.itemstore.band

商品和多个门店关系绑定接口
*/
type TaobaoPlaceStoreItemstoreBandAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 门店id
    _storeIds   []int64
    // 操作类型
    _actionType   string
}

// 初始化TaobaoPlaceStoreItemstoreBandAPIRequest对象
func NewTaobaoPlaceStoreItemstoreBandRequest() *TaobaoPlaceStoreItemstoreBandAPIRequest{
    return &TaobaoPlaceStoreItemstoreBandAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetApiMethodName() string {
    return "taobao.place.store.itemstore.band"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetItemId() int64 {
    return r._itemId
}
// StoreIds Setter
// 门店id
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) SetStoreIds(_storeIds []int64) error {
    r._storeIds = _storeIds
    r.Set("store_ids", _storeIds)
    return nil
}

// StoreIds Getter
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetStoreIds() []int64 {
    return r._storeIds
}
// ActionType Setter
// 操作类型
func (r *TaobaoPlaceStoreItemstoreBandAPIRequest) SetActionType(_actionType string) error {
    r._actionType = _actionType
    r.Set("action_type", _actionType)
    return nil
}

// ActionType Getter
func (r TaobaoPlaceStoreItemstoreBandAPIRequest) GetActionType() string {
    return r._actionType
}