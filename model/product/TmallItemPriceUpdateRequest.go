package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU价格更新接口 API请求
tmall.item.price.update

天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
*/
type TmallItemPriceUpdateRequest struct {
    model.Params
    // 商品ID
    itemId   int64
    // 被更新商品价格
    itemPrice   float64
    // 更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
    skuPrices   []UpdateSkuPrice
    // 商品价格更新时候的可选参数
    options   *UpdateItemPriceOption
}

// 初始化TmallItemPriceUpdateRequest对象
func NewTmallItemPriceUpdateRequest() *TmallItemPriceUpdateRequest{
    return &TmallItemPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemPriceUpdateRequest) GetApiMethodName() string {
    return "tmall.item.price.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品ID
func (r *TmallItemPriceUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallItemPriceUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// ItemPrice Setter
// 被更新商品价格
func (r *TmallItemPriceUpdateRequest) SetItemPrice(itemPrice float64) error {
    r.itemPrice = itemPrice
    r.Set("item_price", itemPrice)
    return nil
}

// ItemPrice Getter
func (r TmallItemPriceUpdateRequest) GetItemPrice() float64 {
    return r.itemPrice
}
// SkuPrices Setter
// 更新SKU价格时候的SKU价格对象；如果没有SKU或者不更新SKU价格，可以不填;查找SKU目前支持ID，属性串和商家编码三种模式，建议选用一种最合适的，切勿滥用，一次调用中如果混合使用，更新结果不可预期！
func (r *TmallItemPriceUpdateRequest) SetSkuPrices(skuPrices []UpdateSkuPrice) error {
    r.skuPrices = skuPrices
    r.Set("sku_prices", skuPrices)
    return nil
}

// SkuPrices Getter
func (r TmallItemPriceUpdateRequest) GetSkuPrices() []UpdateSkuPrice {
    return r.skuPrices
}
// Options Setter
// 商品价格更新时候的可选参数
func (r *TmallItemPriceUpdateRequest) SetOptions(options *UpdateItemPriceOption) error {
    r.options = options
    r.Set("options", options)
    return nil
}

// Options Getter
func (r TmallItemPriceUpdateRequest) GetOptions() *UpdateItemPriceOption {
    return r.options
}
