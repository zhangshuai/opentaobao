package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundCsapplyAPIResponse 商家代客售后提交逆向申请 API返回值
// alibaba.tcls.aelophy.refund.csapply
//
// 商家代客售后提交逆向申请
type AlibabaTclsAelophyRefundCsapplyAPIResponse struct {
	model.CommonResponse
	AlibabaTclsAelophyRefundCsapplyAPIResponseModel
}

// AlibabaTclsAelophyRefundCsapplyAPIResponseModel is 商家代客售后提交逆向申请 成功返回结果
type AlibabaTclsAelophyRefundCsapplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tcls_aelophy_refund_csapply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	ApiResult *AlibabaTclsAelophyRefundCsapplyApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
