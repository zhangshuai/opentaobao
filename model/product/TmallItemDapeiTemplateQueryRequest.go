package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搭配查询接口 API请求
tmall.item.dapei.template.query

根据条件获取搭配内容
*/
type TmallItemDapeiTemplateQueryRequest struct {
    model.Params
    // 搭配标题
    title   string
    // 页码
    pageIndex   int64
    // 分页大小
    pageSize   int64
}

// 初始化TmallItemDapeiTemplateQueryRequest对象
func NewTmallItemDapeiTemplateQueryRequest() *TmallItemDapeiTemplateQueryRequest{
    return &TmallItemDapeiTemplateQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemDapeiTemplateQueryRequest) GetApiMethodName() string {
    return "tmall.item.dapei.template.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemDapeiTemplateQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Title Setter
// 搭配标题
func (r *TmallItemDapeiTemplateQueryRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TmallItemDapeiTemplateQueryRequest) GetTitle() string {
    return r.title
}
// PageIndex Setter
// 页码
func (r *TmallItemDapeiTemplateQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TmallItemDapeiTemplateQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// PageSize Setter
// 分页大小
func (r *TmallItemDapeiTemplateQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TmallItemDapeiTemplateQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
