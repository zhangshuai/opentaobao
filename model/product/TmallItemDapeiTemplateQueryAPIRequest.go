package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallitemdapeitemplatequeryAPIRequest 搭配查询接口 API请求
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
type TmallitemdapeitemplatequeryAPIRequest struct {
	model.Params
	// 搭配标题
	_title string
	// 页码
	_pageIndex int64
	// 分页大小
	_pageSize int64
}

// NewTmallitemdapeitemplatequeryRequest 初始化TmallitemdapeitemplatequeryAPIRequest对象
func NewTmallitemdapeitemplatequeryRequest() *TmallitemdapeitemplatequeryAPIRequest {
	return &TmallitemdapeitemplatequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallitemdapeitemplatequeryAPIRequest) GetApiMethodName() string {
	return "tmall.item.dapei.template.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallitemdapeitemplatequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallitemdapeitemplatequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 搭配标题
func (r *TmallitemdapeitemplatequeryAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TmallitemdapeitemplatequeryAPIRequest) GetTitle() string {
	return r._title
}

// SetPageIndex is PageIndex Setter
// 页码
func (r *TmallitemdapeitemplatequeryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TmallitemdapeitemplatequeryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TmallitemdapeitemplatequeryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallitemdapeitemplatequeryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
