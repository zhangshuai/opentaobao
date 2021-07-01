package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁商品四级车型信息 API请求
tmall.car.lease.itemcarinfo

整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
*/
type TmallCarLeaseItemcarinfoAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
}

// 初始化TmallCarLeaseItemcarinfoAPIRequest对象
func NewTmallCarLeaseItemcarinfoRequest() *TmallCarLeaseItemcarinfoAPIRequest{
    return &TmallCarLeaseItemcarinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseItemcarinfoAPIRequest) GetApiMethodName() string {
    return "tmall.car.lease.itemcarinfo"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseItemcarinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallCarLeaseItemcarinfoAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallCarLeaseItemcarinfoAPIRequest) GetItemId() int64 {
    return r._itemId
}