package uscesl

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按商家批量写入商品接口 API请求
taobao.uscesl.iteminfo.batch.insert

【电子价签】支持按照商家-门店维度批量写入商品数据
*/
type TaobaoUsceslIteminfoBatchInsertAPIRequest struct {
    model.Params
    // 商品信息集合，限制500
    _itemList   []ItemChangeBo
    // 门店ID
    _storeId   int64
    // 商家KEY
    _bizBrandKey   string
}

// 初始化TaobaoUsceslIteminfoBatchInsertAPIRequest对象
func NewTaobaoUsceslIteminfoBatchInsertRequest() *TaobaoUsceslIteminfoBatchInsertAPIRequest{
    return &TaobaoUsceslIteminfoBatchInsertAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetApiMethodName() string {
    return "taobao.uscesl.iteminfo.batch.insert"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemList Setter
// 商品信息集合，限制500
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) SetItemList(_itemList []ItemChangeBo) error {
    r._itemList = _itemList
    r.Set("item_list", _itemList)
    return nil
}

// ItemList Getter
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetItemList() []ItemChangeBo {
    return r._itemList
}
// StoreId Setter
// 门店ID
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetStoreId() int64 {
    return r._storeId
}
// BizBrandKey Setter
// 商家KEY
func (r *TaobaoUsceslIteminfoBatchInsertAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
    r._bizBrandKey = _bizBrandKey
    r.Set("biz_brand_key", _bizBrandKey)
    return nil
}

// BizBrandKey Getter
func (r TaobaoUsceslIteminfoBatchInsertAPIRequest) GetBizBrandKey() string {
    return r._bizBrandKey
}