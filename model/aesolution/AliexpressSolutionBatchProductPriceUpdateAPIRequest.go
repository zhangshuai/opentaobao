package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionBatchProductPriceUpdateAPIRequest aliexpress.solution.batch.product.price.update API请求
// aliexpress.solution.batch.product.price.update
//
// batch product price update operation for oversea sellers
type AliexpressSolutionBatchProductPriceUpdateAPIRequest struct {
	model.Params
	// The product list, in which the price needs to be updated. Maximum length:20
	_mutipleProductUpdateList []SynchronizeProductRequestDto
}

// NewAliexpressSolutionBatchProductPriceUpdateRequest 初始化AliexpressSolutionBatchProductPriceUpdateAPIRequest对象
func NewAliexpressSolutionBatchProductPriceUpdateRequest() *AliexpressSolutionBatchProductPriceUpdateAPIRequest {
	return &AliexpressSolutionBatchProductPriceUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionBatchProductPriceUpdateAPIRequest) Reset() {
	r._mutipleProductUpdateList = r._mutipleProductUpdateList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.batch.product.price.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMutipleProductUpdateList is MutipleProductUpdateList Setter
// The product list, in which the price needs to be updated. Maximum length:20
func (r *AliexpressSolutionBatchProductPriceUpdateAPIRequest) SetMutipleProductUpdateList(_mutipleProductUpdateList []SynchronizeProductRequestDto) error {
	r._mutipleProductUpdateList = _mutipleProductUpdateList
	r.Set("mutiple_product_update_list", _mutipleProductUpdateList)
	return nil
}

// GetMutipleProductUpdateList MutipleProductUpdateList Getter
func (r AliexpressSolutionBatchProductPriceUpdateAPIRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
	return r._mutipleProductUpdateList
}

var poolAliexpressSolutionBatchProductPriceUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionBatchProductPriceUpdateRequest()
	},
}

// GetAliexpressSolutionBatchProductPriceUpdateRequest 从 sync.Pool 获取 AliexpressSolutionBatchProductPriceUpdateAPIRequest
func GetAliexpressSolutionBatchProductPriceUpdateAPIRequest() *AliexpressSolutionBatchProductPriceUpdateAPIRequest {
	return poolAliexpressSolutionBatchProductPriceUpdateAPIRequest.Get().(*AliexpressSolutionBatchProductPriceUpdateAPIRequest)
}

// ReleaseAliexpressSolutionBatchProductPriceUpdateAPIRequest 将 AliexpressSolutionBatchProductPriceUpdateAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionBatchProductPriceUpdateAPIRequest(v *AliexpressSolutionBatchProductPriceUpdateAPIRequest) {
	v.Reset()
	poolAliexpressSolutionBatchProductPriceUpdateAPIRequest.Put(v)
}
