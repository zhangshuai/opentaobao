package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest 查询某计划离线列表 API请求
// taobao.onebp.dkx.report.report.campaign.offline
//
// 查询某计划离线列表；
// 拓展流量查询：
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"launch_product_id_list":[101004013],"start_time":"2021-04-26","campaign_id_list":[134821085],"end_time":"2021-04-28","effect":15,}
// 非拓展流量查询：
// 入参2示例：{"start_time":"2021-09-08","campaign_id_list":[2821811599],"end_time":"2021-09-08","effect":15}
type TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest struct {
	model.Params
	// 请求体
	_apiServiceContext *ApiServiceContext
	// 报表查询参数
	_reportQuery *ReportQueryTopDto
}

// NewTaobaoOnebpDkxReportReportCampaignOfflineRequest 初始化TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest对象
func NewTaobaoOnebpDkxReportReportCampaignOfflineRequest() *TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest {
	return &TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.report.report.campaign.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetApiServiceContext is ApiServiceContext Setter
// 请求体
func (r *TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) SetApiServiceContext(_apiServiceContext *ApiServiceContext) error {
	r._apiServiceContext = _apiServiceContext
	r.Set("api_service_context", _apiServiceContext)
	return nil
}

// GetApiServiceContext ApiServiceContext Getter
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetApiServiceContext() *ApiServiceContext {
	return r._apiServiceContext
}

// SetReportQuery is ReportQuery Setter
// 报表查询参数
func (r *TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) SetReportQuery(_reportQuery *ReportQueryTopDto) error {
	r._reportQuery = _reportQuery
	r.Set("report_query", _reportQuery)
	return nil
}

// GetReportQuery ReportQuery Getter
func (r TaobaoOnebpDkxReportReportCampaignOfflineAPIRequest) GetReportQuery() *ReportQueryTopDto {
	return r._reportQuery
}