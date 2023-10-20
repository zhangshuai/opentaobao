package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmpointquerypointflowAPIResponse 分页查询积分流水 API返回值
// alibaba.alsc.crm.point.querypointflow
//
// 分页查询积分流水
type AlibabaalsccrmpointquerypointflowAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmpointquerypointflowAPIResponseModel
}

// AlibabaalsccrmpointquerypointflowAPIResponseModel is 分页查询积分流水 成功返回结果
type AlibabaalsccrmpointquerypointflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_querypointflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
