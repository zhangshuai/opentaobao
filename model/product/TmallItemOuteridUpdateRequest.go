package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU商家编码更新接口 API请求
tmall.item.outerid.update

天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
*/
type TmallItemOuteridUpdateRequest struct {
    model.Params
    // 商品ID
    itemId   int64
    // 商品维度商家编码，如果不修改可以不传；清空请设置空串
    outerId   string
    // 商品SKU更新OuterId时候用的数据
    skuOuters   []UpdateSkuOuterId
}

// 初始化TmallItemOuteridUpdateRequest对象
func NewTmallItemOuteridUpdateRequest() *TmallItemOuteridUpdateRequest{
    return &TmallItemOuteridUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemOuteridUpdateRequest) GetApiMethodName() string {
    return "tmall.item.outerid.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemOuteridUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TmallItemOuteridUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallItemOuteridUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// OuterId Setter
// 商品维度商家编码，如果不修改可以不传；清空请设置空串
func (r *TmallItemOuteridUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TmallItemOuteridUpdateRequest) GetOuterId() string {
    return r.outerId
}
// SkuOuters Setter
// 商品SKU更新OuterId时候用的数据
func (r *TmallItemOuteridUpdateRequest) SetSkuOuters(skuOuters []UpdateSkuOuterId) error {
    r.skuOuters = skuOuters
    r.Set("sku_outers", skuOuters)
    return nil
}

// SkuOuters Getter
func (r TmallItemOuteridUpdateRequest) GetSkuOuters() []UpdateSkuOuterId {
    return r.skuOuters
}
