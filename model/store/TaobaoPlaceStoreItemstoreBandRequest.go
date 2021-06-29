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
type TaobaoPlaceStoreItemstoreBandRequest struct {
    model.Params
    // 商品id
    itemId   int64
    // 门店id
    storeIds   []int64
    // 操作类型
    actionType   string
}

// 初始化TaobaoPlaceStoreItemstoreBandRequest对象
func NewTaobaoPlaceStoreItemstoreBandRequest() *TaobaoPlaceStoreItemstoreBandRequest{
    return &TaobaoPlaceStoreItemstoreBandRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStoreItemstoreBandRequest) GetApiMethodName() string {
    return "taobao.place.store.itemstore.band"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStoreItemstoreBandRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoPlaceStoreItemstoreBandRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoPlaceStoreItemstoreBandRequest) GetItemId() int64 {
    return r.itemId
}
// StoreIds Setter
// 门店id
func (r *TaobaoPlaceStoreItemstoreBandRequest) SetStoreIds(storeIds []int64) error {
    r.storeIds = storeIds
    r.Set("store_ids", storeIds)
    return nil
}

// StoreIds Getter
func (r TaobaoPlaceStoreItemstoreBandRequest) GetStoreIds() []int64 {
    return r.storeIds
}
// ActionType Setter
// 操作类型
func (r *TaobaoPlaceStoreItemstoreBandRequest) SetActionType(actionType string) error {
    r.actionType = actionType
    r.Set("action_type", actionType)
    return nil
}

// ActionType Getter
func (r TaobaoPlaceStoreItemstoreBandRequest) GetActionType() string {
    return r.actionType
}
