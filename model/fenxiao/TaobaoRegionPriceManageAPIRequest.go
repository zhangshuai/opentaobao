package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑区域价格 API请求
taobao.region.price.manage

编辑区域价格
*/
type TaobaoRegionPriceManageAPIRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 无sku传0
    _skuId   int64
    // 列表
    _regionalPriceDtos   []RegionalPriceDto
    // true:全量, false:增量
    _isFull   bool
}

// 初始化TaobaoRegionPriceManageAPIRequest对象
func NewTaobaoRegionPriceManageRequest() *TaobaoRegionPriceManageAPIRequest{
    return &TaobaoRegionPriceManageAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRegionPriceManageAPIRequest) GetApiMethodName() string {
    return "taobao.region.price.manage"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRegionPriceManageAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TaobaoRegionPriceManageAPIRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoRegionPriceManageAPIRequest) GetItemId() int64 {
    return r._itemId
}
// SkuId Setter
// 无sku传0
func (r *TaobaoRegionPriceManageAPIRequest) SetSkuId(_skuId int64) error {
    r._skuId = _skuId
    r.Set("sku_id", _skuId)
    return nil
}

// SkuId Getter
func (r TaobaoRegionPriceManageAPIRequest) GetSkuId() int64 {
    return r._skuId
}
// RegionalPriceDtos Setter
// 列表
func (r *TaobaoRegionPriceManageAPIRequest) SetRegionalPriceDtos(_regionalPriceDtos []RegionalPriceDto) error {
    r._regionalPriceDtos = _regionalPriceDtos
    r.Set("regional_price_dtos", _regionalPriceDtos)
    return nil
}

// RegionalPriceDtos Getter
func (r TaobaoRegionPriceManageAPIRequest) GetRegionalPriceDtos() []RegionalPriceDto {
    return r._regionalPriceDtos
}
// IsFull Setter
// true:全量, false:增量
func (r *TaobaoRegionPriceManageAPIRequest) SetIsFull(_isFull bool) error {
    r._isFull = _isFull
    r.Set("is_full", _isFull)
    return nil
}

// IsFull Getter
func (r TaobaoRegionPriceManageAPIRequest) GetIsFull() bool {
    return r._isFull
}