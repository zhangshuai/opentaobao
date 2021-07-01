package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置P4P产品优先推广状态 API请求
alibaba.scbp.product.preferential.update

设置P4P产品优先推广状态
*/
type AlibabaScbpProductPreferentialUpdateAPIRequest struct {
    model.Params
    // 关键词ID
    _keywordId   int64
    // 产品ID
    _productId   int64
    // Y:设置优推,N:取消优推
    _status   string
}

// 初始化AlibabaScbpProductPreferentialUpdateAPIRequest对象
func NewAlibabaScbpProductPreferentialUpdateRequest() *AlibabaScbpProductPreferentialUpdateAPIRequest{
    return &AlibabaScbpProductPreferentialUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.product.preferential.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// KeywordId Setter
// 关键词ID
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) SetKeywordId(_keywordId int64) error {
    r._keywordId = _keywordId
    r.Set("keyword_id", _keywordId)
    return nil
}

// KeywordId Getter
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetKeywordId() int64 {
    return r._keywordId
}
// ProductId Setter
// 产品ID
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) SetProductId(_productId int64) error {
    r._productId = _productId
    r.Set("product_id", _productId)
    return nil
}

// ProductId Getter
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetProductId() int64 {
    return r._productId
}
// Status Setter
// Y:设置优推,N:取消优推
func (r *AlibabaScbpProductPreferentialUpdateAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaScbpProductPreferentialUpdateAPIRequest) GetStatus() string {
    return r._status
}