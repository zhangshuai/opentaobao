package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse 二级物料-LT内的POGAP数据回传 API返回值
// alibaba.tmallgenie.scp.plan.rawpo.gap.return
//
// 二级物料-LT内的POGAP数据回传
type AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponseModel is 二级物料-LT内的POGAP数据回传 成功返回结果
type AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_rawpo_gap_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse
func GetAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse() *AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse {
	return poolAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse.Get().(*AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse 将 AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse(v *AlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanRawpoGapReturnAPIResponse.Put(v)
}
