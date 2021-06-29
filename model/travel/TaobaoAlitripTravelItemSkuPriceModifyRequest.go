package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】日期级别日历价格库存修改，增量维护 API请求
taobao.alitrip.travel.item.sku.price.modify

【API3.0】日期级别日历价格库存增量维护
*/
type TaobaoAlitripTravelItemSkuPriceModifyRequest struct {
    model.Params
    // 商品id。itemId和outProductId至少填写一个
    itemId   int64
    // 商品 外部商家编码。itemId和outProductId至少填写一个
    outProductId   string
    // 商品日历价格库存套餐
    skus   []PontusTravelItemSkuInfo
}

// 初始化TaobaoAlitripTravelItemSkuPriceModifyRequest对象
func NewTaobaoAlitripTravelItemSkuPriceModifyRequest() *TaobaoAlitripTravelItemSkuPriceModifyRequest{
    return &TaobaoAlitripTravelItemSkuPriceModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemSkuPriceModifyRequest) GetApiMethodName() string {
    return "taobao.alitrip.travel.item.sku.price.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemSkuPriceModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyRequest) GetItemId() int64 {
    return r.itemId
}
// OutProductId Setter
// 商品 外部商家编码。itemId和outProductId至少填写一个
func (r *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyRequest) GetOutProductId() string {
    return r.outProductId
}
// Skus Setter
// 商品日历价格库存套餐
func (r *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetSkus(skus []PontusTravelItemSkuInfo) error {
    r.skus = skus
    r.Set("skus", skus)
    return nil
}

// Skus Getter
func (r TaobaoAlitripTravelItemSkuPriceModifyRequest) GetSkus() []PontusTravelItemSkuInfo {
    return r.skus
}
