package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCampaignAPIResponse 计划报表查询 API返回值
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
type TaobaoUniversalbpReportQueryCampaignAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryCampaignAPIResponseModel
}

// TaobaoUniversalbpReportQueryCampaignAPIResponseModel is 计划报表查询 成功返回结果
type TaobaoUniversalbpReportQueryCampaignAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_campaign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryCampaignTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
