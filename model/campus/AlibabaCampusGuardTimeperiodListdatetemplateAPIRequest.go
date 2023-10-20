package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardtimeperiodlistdatetemplateAPIRequest 门禁控制器查询日期模版 API请求
// alibaba.campus.guard.timeperiod.listdatetemplate
//
// 门禁控制器查询日期模版
type AlibabacampusguardtimeperiodlistdatetemplateAPIRequest struct {
	model.Params
	// 参数
	_calenderTemplateQuery *CalenderTemplateQuery
}

// NewAlibabacampusguardtimeperiodlistdatetemplateRequest 初始化AlibabacampusguardtimeperiodlistdatetemplateAPIRequest对象
func NewAlibabacampusguardtimeperiodlistdatetemplateRequest() *AlibabacampusguardtimeperiodlistdatetemplateAPIRequest {
	return &AlibabacampusguardtimeperiodlistdatetemplateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusguardtimeperiodlistdatetemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.guard.timeperiod.listdatetemplate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusguardtimeperiodlistdatetemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusguardtimeperiodlistdatetemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCalenderTemplateQuery is CalenderTemplateQuery Setter
// 参数
func (r *AlibabacampusguardtimeperiodlistdatetemplateAPIRequest) SetCalenderTemplateQuery(_calenderTemplateQuery *CalenderTemplateQuery) error {
	r._calenderTemplateQuery = _calenderTemplateQuery
	r.Set("calender_template_query", _calenderTemplateQuery)
	return nil
}

// GetCalenderTemplateQuery CalenderTemplateQuery Getter
func (r AlibabacampusguardtimeperiodlistdatetemplateAPIRequest) GetCalenderTemplateQuery() *CalenderTemplateQuery {
	return r._calenderTemplateQuery
}