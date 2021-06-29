package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.batch.product.inventory.update API请求
aliexpress.solution.batch.product.inventory.update

batch product inventory update API for oversea sellers. Sellers could update multiple skus among multiple products in a single call. Maximum 20 products could be updated at the same time and maximum 200 skus could be updated within one product.
*/
type AliexpressSolutionBatchProductInventoryUpdateRequest struct {
    model.Params
    // The product list, in which the inventory needs to be updated. Maximum 20 products.
    mutipleProductUpdateList   []SynchronizeProductRequestDto
}

// 初始化AliexpressSolutionBatchProductInventoryUpdateRequest对象
func NewAliexpressSolutionBatchProductInventoryUpdateRequest() *AliexpressSolutionBatchProductInventoryUpdateRequest{
    return &AliexpressSolutionBatchProductInventoryUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionBatchProductInventoryUpdateRequest) GetApiMethodName() string {
    return "aliexpress.solution.batch.product.inventory.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionBatchProductInventoryUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MutipleProductUpdateList Setter
// The product list, in which the inventory needs to be updated. Maximum 20 products.
func (r *AliexpressSolutionBatchProductInventoryUpdateRequest) SetMutipleProductUpdateList(mutipleProductUpdateList []SynchronizeProductRequestDto) error {
    r.mutipleProductUpdateList = mutipleProductUpdateList
    r.Set("mutiple_product_update_list", mutipleProductUpdateList)
    return nil
}

// MutipleProductUpdateList Getter
func (r AliexpressSolutionBatchProductInventoryUpdateRequest) GetMutipleProductUpdateList() []SynchronizeProductRequestDto {
    return r.mutipleProductUpdateList
}
