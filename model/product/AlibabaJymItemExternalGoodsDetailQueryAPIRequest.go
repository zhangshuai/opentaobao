package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsDetailQueryAPIRequest 交易猫外部商家商品详情查询接口 API请求
// alibaba.jym.item.external.goods.detail.query
//
// 供外部B端商家接入，请求查询商品详情，返回商品详情查询结果
type AlibabaJymItemExternalGoodsDetailQueryAPIRequest struct {
	model.Params
	// 商品详情查询接口入参
	_goodsDetailQuery *GoodsDetailQueryDto
}

// NewAlibabaJymItemExternalGoodsDetailQueryRequest 初始化AlibabaJymItemExternalGoodsDetailQueryAPIRequest对象
func NewAlibabaJymItemExternalGoodsDetailQueryRequest() *AlibabaJymItemExternalGoodsDetailQueryAPIRequest {
	return &AlibabaJymItemExternalGoodsDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsDetailQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDetailQuery is GoodsDetailQuery Setter
// 商品详情查询接口入参
func (r *AlibabaJymItemExternalGoodsDetailQueryAPIRequest) SetGoodsDetailQuery(_goodsDetailQuery *GoodsDetailQueryDto) error {
	r._goodsDetailQuery = _goodsDetailQuery
	r.Set("goods_detail_query", _goodsDetailQuery)
	return nil
}

// GetGoodsDetailQuery GoodsDetailQuery Getter
func (r AlibabaJymItemExternalGoodsDetailQueryAPIRequest) GetGoodsDetailQuery() *GoodsDetailQueryDto {
	return r._goodsDetailQuery
}
