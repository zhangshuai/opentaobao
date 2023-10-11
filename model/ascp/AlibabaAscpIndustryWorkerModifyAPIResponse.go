package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryWorkerModifyAPIResponse 送货入户并安装修改师傅信息 API返回值
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
type AlibabaAscpIndustryWorkerModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryWorkerModifyAPIResponseModel
}

// AlibabaAscpIndustryWorkerModifyAPIResponseModel is 送货入户并安装修改师傅信息 成功返回结果
type AlibabaAscpIndustryWorkerModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_worker_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
