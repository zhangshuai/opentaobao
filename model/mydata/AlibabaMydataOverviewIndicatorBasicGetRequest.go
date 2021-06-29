package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取公司询盘流量行业表现 API请求
alibaba.mydata.overview.indicator.basic.get

获取公司询盘流量行业表现
*/
type AlibabaMydataOverviewIndicatorBasicGetRequest struct {
    model.Params
    // 要查询的数据周期
    dateRange   *DateRange
    // 要查询的行业信息
    industry   *Industry
}

// 初始化AlibabaMydataOverviewIndicatorBasicGetRequest对象
func NewAlibabaMydataOverviewIndicatorBasicGetRequest() *AlibabaMydataOverviewIndicatorBasicGetRequest{
    return &AlibabaMydataOverviewIndicatorBasicGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewIndicatorBasicGetRequest) GetApiMethodName() string {
    return "alibaba.mydata.overview.indicator.basic.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewIndicatorBasicGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DateRange Setter
// 要查询的数据周期
func (r *AlibabaMydataOverviewIndicatorBasicGetRequest) SetDateRange(dateRange *DateRange) error {
    r.dateRange = dateRange
    r.Set("date_range", dateRange)
    return nil
}

// DateRange Getter
func (r AlibabaMydataOverviewIndicatorBasicGetRequest) GetDateRange() *DateRange {
    return r.dateRange
}
// Industry Setter
// 要查询的行业信息
func (r *AlibabaMydataOverviewIndicatorBasicGetRequest) SetIndustry(industry *Industry) error {
    r.industry = industry
    r.Set("industry", industry)
    return nil
}

// Industry Getter
func (r AlibabaMydataOverviewIndicatorBasicGetRequest) GetIndustry() *Industry {
    return r.industry
}