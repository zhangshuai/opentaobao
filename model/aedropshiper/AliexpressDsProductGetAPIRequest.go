package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsProductGetAPIRequest 商品信息查询 API请求
// aliexpress.ds.product.get
//
// 商品信息查询
type AliexpressDsProductGetAPIRequest struct {
	model.Params
	// Country
	_shipToCountry string
	// Target currency
	_targetCurrency string
	// Target language
	_targetLanguage string
	// Item ID
	_productId int64
}

// NewAliexpressDsProductGetRequest 初始化AliexpressDsProductGetAPIRequest对象
func NewAliexpressDsProductGetRequest() *AliexpressDsProductGetAPIRequest {
	return &AliexpressDsProductGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressDsProductGetAPIRequest) Reset() {
	r._shipToCountry = ""
	r._targetCurrency = ""
	r._targetLanguage = ""
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressDsProductGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.product.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressDsProductGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressDsProductGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShipToCountry is ShipToCountry Setter
// Country
func (r *AliexpressDsProductGetAPIRequest) SetShipToCountry(_shipToCountry string) error {
	r._shipToCountry = _shipToCountry
	r.Set("ship_to_country", _shipToCountry)
	return nil
}

// GetShipToCountry ShipToCountry Getter
func (r AliexpressDsProductGetAPIRequest) GetShipToCountry() string {
	return r._shipToCountry
}

// SetTargetCurrency is TargetCurrency Setter
// Target currency
func (r *AliexpressDsProductGetAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressDsProductGetAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// Target language
func (r *AliexpressDsProductGetAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressDsProductGetAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetProductId is ProductId Setter
// Item ID
func (r *AliexpressDsProductGetAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AliexpressDsProductGetAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAliexpressDsProductGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressDsProductGetRequest()
	},
}

// GetAliexpressDsProductGetRequest 从 sync.Pool 获取 AliexpressDsProductGetAPIRequest
func GetAliexpressDsProductGetAPIRequest() *AliexpressDsProductGetAPIRequest {
	return poolAliexpressDsProductGetAPIRequest.Get().(*AliexpressDsProductGetAPIRequest)
}

// ReleaseAliexpressDsProductGetAPIRequest 将 AliexpressDsProductGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressDsProductGetAPIRequest(v *AliexpressDsProductGetAPIRequest) {
	v.Reset()
	poolAliexpressDsProductGetAPIRequest.Put(v)
}
