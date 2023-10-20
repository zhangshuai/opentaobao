package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductbatchupdatedisplayAPIRequest 商品批量上下架接口 API请求
// alibaba.icbu.product.batch.update.display
//
// 给国际站的三方服务商提供批量上下架接口
type AlibabaicbuproductbatchupdatedisplayAPIRequest struct {
	model.Params
	// on表示上架，off表示下架
	_newDisplay string
	// 用逗号分隔的混淆id字符串
	_productIdList string
}

// NewAlibabaicbuproductbatchupdatedisplayRequest 初始化AlibabaicbuproductbatchupdatedisplayAPIRequest对象
func NewAlibabaicbuproductbatchupdatedisplayRequest() *AlibabaicbuproductbatchupdatedisplayAPIRequest {
	return &AlibabaicbuproductbatchupdatedisplayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductbatchupdatedisplayAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.batch.update.display"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductbatchupdatedisplayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductbatchupdatedisplayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNewDisplay is NewDisplay Setter
// on表示上架，off表示下架
func (r *AlibabaicbuproductbatchupdatedisplayAPIRequest) SetNewDisplay(_newDisplay string) error {
	r._newDisplay = _newDisplay
	r.Set("new_display", _newDisplay)
	return nil
}

// GetNewDisplay NewDisplay Getter
func (r AlibabaicbuproductbatchupdatedisplayAPIRequest) GetNewDisplay() string {
	return r._newDisplay
}

// SetProductIdList is ProductIdList Setter
// 用逗号分隔的混淆id字符串
func (r *AlibabaicbuproductbatchupdatedisplayAPIRequest) SetProductIdList(_productIdList string) error {
	r._productIdList = _productIdList
	r.Set("product_id_list", _productIdList)
	return nil
}

// GetProductIdList ProductIdList Getter
func (r AlibabaicbuproductbatchupdatedisplayAPIRequest) GetProductIdList() string {
	return r._productIdList
}
