package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询盒马帮档期活动详情 API请求
alibaba.price.promotion.activity.query

查询盒马帮档期活动详情
*/
type AlibabaPricePromotionActivityQueryAPIRequest struct {
    model.Params
    // 页码
    _page   int64
    // 外部档期code
    _outerPromotionCode   string
    // TOB店仓编码
    _ouCode   string
    // 页码大小
    _pageSize   int64
}

// 初始化AlibabaPricePromotionActivityQueryAPIRequest对象
func NewAlibabaPricePromotionActivityQueryRequest() *AlibabaPricePromotionActivityQueryAPIRequest{
    return &AlibabaPricePromotionActivityQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.price.promotion.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Page Setter
// 页码
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetPage() int64 {
    return r._page
}
// OuterPromotionCode Setter
// 外部档期code
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetOuterPromotionCode(_outerPromotionCode string) error {
    r._outerPromotionCode = _outerPromotionCode
    r.Set("outer_promotion_code", _outerPromotionCode)
    return nil
}

// OuterPromotionCode Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetOuterPromotionCode() string {
    return r._outerPromotionCode
}
// OuCode Setter
// TOB店仓编码
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetOuCode(_ouCode string) error {
    r._ouCode = _ouCode
    r.Set("ou_code", _ouCode)
    return nil
}

// OuCode Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetOuCode() string {
    return r._ouCode
}
// PageSize Setter
// 页码大小
func (r *AlibabaPricePromotionActivityQueryAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaPricePromotionActivityQueryAPIRequest) GetPageSize() int64 {
    return r._pageSize
}