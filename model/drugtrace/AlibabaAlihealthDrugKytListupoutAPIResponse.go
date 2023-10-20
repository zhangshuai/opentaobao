package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytlistupoutAPIResponse 查询货主/本企业上游企业出库单据信息 API返回值
// alibaba.alihealth.drug.kyt.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaalihealthdrugkytlistupoutAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytlistupoutAPIResponseModel
}

// AlibabaalihealthdrugkytlistupoutAPIResponseModel is 查询货主/本企业上游企业出库单据信息 成功返回结果
type AlibabaalihealthdrugkytlistupoutAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_listupout_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytlistupoutResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
