package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划维度小时报表获取 API返回值 
taobao.simba.hour.report.campaign.get

计划维度小时报表获取
*/
type TaobaoSimbaHourReportCampaignGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaHourReportCampaignGetResponse
}

// 计划维度小时报表获取 成功返回结果
type TaobaoSimbaHourReportCampaignGetResponse struct {
    XMLName xml.Name `xml:"simba_hour_report_campaign_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 11
    Results   *RtRptResultEntityDTO `json:"results,omitempty" xml:"results,omitempty"`
}
