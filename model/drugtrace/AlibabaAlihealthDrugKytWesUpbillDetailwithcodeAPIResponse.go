package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponse 查询上游出库单明细（带追溯码信息） API返回值
// alibaba.alihealth.drug.kyt.wes.upbill.detailwithcode
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponseModel
}

// AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponseModel is 查询上游出库单明细（带追溯码信息） 成功返回结果
type AlibabaalihealthdrugkytwesupbilldetailwithcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_upbill_detailwithcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 监控宝推送网站监控信息，返回结果
	Result *AlibabaalihealthdrugkytwesupbilldetailwithcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}