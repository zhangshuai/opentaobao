package foodscan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabafootscanminiquerymobilereportAPIResponse 根据scanId查询报告 API返回值
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
type AlibabafootscanminiquerymobilereportAPIResponse struct {
	model.CommonResponse
	AlibabafootscanminiquerymobilereportAPIResponseModel
}

// AlibabafootscanminiquerymobilereportAPIResponseModel is 根据scanId查询报告 成功返回结果
type AlibabafootscanminiquerymobilereportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_footscan_mini_query_mobilereport_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务出参
	Result *AlibabafootscanminiquerymobilereportMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
