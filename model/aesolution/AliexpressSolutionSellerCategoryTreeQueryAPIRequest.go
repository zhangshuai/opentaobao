package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.seller.category.tree.query API请求
aliexpress.solution.seller.category.tree.query

API for seller to query the category tree. Support only displaying the categories which seller have permissions to publish products.
*/
type AliexpressSolutionSellerCategoryTreeQueryAPIRequest struct {
    model.Params
    // parent category ID. To obtain the root categories, setting the category_id = 0
    _categoryId   int64
    // filter the categories which seller does not have permission
    _filterNoPermission   bool
}

// 初始化AliexpressSolutionSellerCategoryTreeQueryAPIRequest对象
func NewAliexpressSolutionSellerCategoryTreeQueryRequest() *AliexpressSolutionSellerCategoryTreeQueryAPIRequest{
    return &AliexpressSolutionSellerCategoryTreeQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.seller.category.tree.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CategoryId Setter
// parent category ID. To obtain the root categories, setting the category_id = 0
func (r *AliexpressSolutionSellerCategoryTreeQueryAPIRequest) SetCategoryId(_categoryId int64) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetCategoryId() int64 {
    return r._categoryId
}
// FilterNoPermission Setter
// filter the categories which seller does not have permission
func (r *AliexpressSolutionSellerCategoryTreeQueryAPIRequest) SetFilterNoPermission(_filterNoPermission bool) error {
    r._filterNoPermission = _filterNoPermission
    r.Set("filter_no_permission", _filterNoPermission)
    return nil
}

// FilterNoPermission Getter
func (r AliexpressSolutionSellerCategoryTreeQueryAPIRequest) GetFilterNoPermission() bool {
    return r._filterNoPermission
}