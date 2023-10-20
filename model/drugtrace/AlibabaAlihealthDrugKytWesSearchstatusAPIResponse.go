package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwessearchstatusAPIResponse 单据处理状态查询 API返回值
// alibaba.alihealth.drug.kyt.wes.searchstatus
//
// 单据处理状态查询
type AlibabaalihealthdrugkytwessearchstatusAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytwessearchstatusAPIResponseModel
}

// AlibabaalihealthdrugkytwessearchstatusAPIResponseModel is 单据处理状态查询 成功返回结果
type AlibabaalihealthdrugkytwessearchstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_searchstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytwessearchstatusResultModel `json:"result,omitempty" xml:"result,omitempty"`
}