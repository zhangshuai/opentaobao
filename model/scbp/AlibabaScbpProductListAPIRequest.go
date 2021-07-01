package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询P4P产品 API请求
alibaba.scbp.product.list

查询P4P产品
*/
type AlibabaScbpProductListAPIRequest struct {
    model.Params
    // 产品分组标识
    _groupId   string
    // 产品分页查询，每页个数，最大值20
    _perPageSize   int64
    // 第几页
    _toPage   int64
}

// 初始化AlibabaScbpProductListAPIRequest对象
func NewAlibabaScbpProductListRequest() *AlibabaScbpProductListAPIRequest{
    return &AlibabaScbpProductListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductListAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 产品分组标识
func (r *AlibabaScbpProductListAPIRequest) SetGroupId(_groupId string) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaScbpProductListAPIRequest) GetGroupId() string {
    return r._groupId
}
// PerPageSize Setter
// 产品分页查询，每页个数，最大值20
func (r *AlibabaScbpProductListAPIRequest) SetPerPageSize(_perPageSize int64) error {
    r._perPageSize = _perPageSize
    r.Set("per_page_size", _perPageSize)
    return nil
}

// PerPageSize Getter
func (r AlibabaScbpProductListAPIRequest) GetPerPageSize() int64 {
    return r._perPageSize
}
// ToPage Setter
// 第几页
func (r *AlibabaScbpProductListAPIRequest) SetToPage(_toPage int64) error {
    r._toPage = _toPage
    r.Set("to_page", _toPage)
    return nil
}

// ToPage Getter
func (r AlibabaScbpProductListAPIRequest) GetToPage() int64 {
    return r._toPage
}