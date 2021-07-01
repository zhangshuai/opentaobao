package store

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店和子门店关系删除 API请求
taobao.place.storerelatesub.delete

门店和子门店关系删除
*/
type TaobaoPlaceStorerelatesubDeleteAPIRequest struct {
    model.Params
    // 门店Id
    _storeId   int64
    // 子门店id
    _subStoreIds   []int64
}

// 初始化TaobaoPlaceStorerelatesubDeleteAPIRequest对象
func NewTaobaoPlaceStorerelatesubDeleteRequest() *TaobaoPlaceStorerelatesubDeleteAPIRequest{
    return &TaobaoPlaceStorerelatesubDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPlaceStorerelatesubDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.place.storerelatesub.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPlaceStorerelatesubDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店Id
func (r *TaobaoPlaceStorerelatesubDeleteAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoPlaceStorerelatesubDeleteAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// SubStoreIds Setter
// 子门店id
func (r *TaobaoPlaceStorerelatesubDeleteAPIRequest) SetSubStoreIds(_subStoreIds []int64) error {
    r._subStoreIds = _subStoreIds
    r.Set("sub_store_ids", _subStoreIds)
    return nil
}

// SubStoreIds Getter
func (r TaobaoPlaceStorerelatesubDeleteAPIRequest) GetSubStoreIds() []int64 {
    return r._subStoreIds
}