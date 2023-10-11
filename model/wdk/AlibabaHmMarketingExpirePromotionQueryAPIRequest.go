package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingExpirePromotionQueryAPIRequest 短保优惠查询 API请求
// alibaba.hm.marketing.expire.promotion.query
//
// 短保优惠查询
type AlibabaHmMarketingExpirePromotionQueryAPIRequest struct {
	model.Params
	// 店铺id
	_shopId string
	// 商品skucode
	_skuCode string
}

// NewAlibabaHmMarketingExpirePromotionQueryRequest 初始化AlibabaHmMarketingExpirePromotionQueryAPIRequest对象
func NewAlibabaHmMarketingExpirePromotionQueryRequest() *AlibabaHmMarketingExpirePromotionQueryAPIRequest {
	return &AlibabaHmMarketingExpirePromotionQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingExpirePromotionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.expire.promotion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingExpirePromotionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingExpirePromotionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 店铺id
func (r *AlibabaHmMarketingExpirePromotionQueryAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaHmMarketingExpirePromotionQueryAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSkuCode is SkuCode Setter
// 商品skucode
func (r *AlibabaHmMarketingExpirePromotionQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaHmMarketingExpirePromotionQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}
