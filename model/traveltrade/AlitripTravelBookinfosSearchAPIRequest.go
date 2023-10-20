package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelbookinfossearchAPIRequest 飞猪度假-订单预定信息列表搜索接口 API请求
// alitrip.travel.bookinfos.search
//
// 查询订单预定信息列表
type AlitriptravelbookinfossearchAPIRequest struct {
	model.Params
	// 申请时间_结束，精确到分钟
	_applyTimeEnd string
	// 申请时间_开始，精确到分钟
	_applyTimeStart string
	// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
	_pageSize int64
	// 当前页
	_currentPage int64
}

// NewAlitriptravelbookinfossearchRequest 初始化AlitriptravelbookinfossearchAPIRequest对象
func NewAlitriptravelbookinfossearchRequest() *AlitriptravelbookinfossearchAPIRequest {
	return &AlitriptravelbookinfossearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelbookinfossearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.bookinfos.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelbookinfossearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelbookinfossearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyTimeEnd is ApplyTimeEnd Setter
// 申请时间_结束，精确到分钟
func (r *AlitriptravelbookinfossearchAPIRequest) SetApplyTimeEnd(_applyTimeEnd string) error {
	r._applyTimeEnd = _applyTimeEnd
	r.Set("apply_time_end", _applyTimeEnd)
	return nil
}

// GetApplyTimeEnd ApplyTimeEnd Getter
func (r AlitriptravelbookinfossearchAPIRequest) GetApplyTimeEnd() string {
	return r._applyTimeEnd
}

// SetApplyTimeStart is ApplyTimeStart Setter
// 申请时间_开始，精确到分钟
func (r *AlitriptravelbookinfossearchAPIRequest) SetApplyTimeStart(_applyTimeStart string) error {
	r._applyTimeStart = _applyTimeStart
	r.Set("apply_time_start", _applyTimeStart)
	return nil
}

// GetApplyTimeStart ApplyTimeStart Getter
func (r AlitriptravelbookinfossearchAPIRequest) GetApplyTimeStart() string {
	return r._applyTimeStart
}

// SetPageSize is PageSize Setter
// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
func (r *AlitriptravelbookinfossearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitriptravelbookinfossearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlitriptravelbookinfossearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlitriptravelbookinfossearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
