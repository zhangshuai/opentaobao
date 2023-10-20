package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwuscoresgetAPIRequest 服务平台评价查询接口 API请求
// taobao.fuwu.scores.get
//
// 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaofuwuscoresgetAPIRequest struct {
	model.Params
	// 评价日期，查询某一天的评价
	_date string
	// 每页获取条数。默认值40，最小值1，最大值100。
	_pageSize int64
	// 当前页
	_currentPage int64
}

// NewTaobaofuwuscoresgetRequest 初始化TaobaofuwuscoresgetAPIRequest对象
func NewTaobaofuwuscoresgetRequest() *TaobaofuwuscoresgetAPIRequest {
	return &TaobaofuwuscoresgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofuwuscoresgetAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.scores.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofuwuscoresgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofuwuscoresgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 评价日期，查询某一天的评价
func (r *TaobaofuwuscoresgetAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r TaobaofuwuscoresgetAPIRequest) GetDate() string {
	return r._date
}

// SetPageSize is PageSize Setter
// 每页获取条数。默认值40，最小值1，最大值100。
func (r *TaobaofuwuscoresgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofuwuscoresgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *TaobaofuwuscoresgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r TaobaofuwuscoresgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
