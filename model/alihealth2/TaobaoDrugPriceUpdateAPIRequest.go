package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家更新宝贝价格 API请求
taobao.drug.price.update

商家更新价格
*/
type TaobaoDrugPriceUpdateAPIRequest struct {
    model.Params
    // 对应的外部店铺ID
    _outStoreId   string
    // 对应的外部商品编码
    _outItemId   string
    // 商品价格
    _price   float64
}

// 初始化TaobaoDrugPriceUpdateAPIRequest对象
func NewTaobaoDrugPriceUpdateRequest() *TaobaoDrugPriceUpdateAPIRequest{
    return &TaobaoDrugPriceUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDrugPriceUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.drug.price.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDrugPriceUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutStoreId Setter
// 对应的外部店铺ID
func (r *TaobaoDrugPriceUpdateAPIRequest) SetOutStoreId(_outStoreId string) error {
    r._outStoreId = _outStoreId
    r.Set("out_store_id", _outStoreId)
    return nil
}

// OutStoreId Getter
func (r TaobaoDrugPriceUpdateAPIRequest) GetOutStoreId() string {
    return r._outStoreId
}
// OutItemId Setter
// 对应的外部商品编码
func (r *TaobaoDrugPriceUpdateAPIRequest) SetOutItemId(_outItemId string) error {
    r._outItemId = _outItemId
    r.Set("out_item_id", _outItemId)
    return nil
}

// OutItemId Getter
func (r TaobaoDrugPriceUpdateAPIRequest) GetOutItemId() string {
    return r._outItemId
}
// Price Setter
// 商品价格
func (r *TaobaoDrugPriceUpdateAPIRequest) SetPrice(_price float64) error {
    r._price = _price
    r.Set("price", _price)
    return nil
}

// Price Getter
func (r TaobaoDrugPriceUpdateAPIRequest) GetPrice() float64 {
    return r._price
}