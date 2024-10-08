package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse 同步销售数据按照渠道类型汇总 API返回值
// alibaba.tmallgenie.scp.plan.summary.sale.qty.get
//
// 同步销售数据按照渠道类型汇总
type AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponseModel is 同步销售数据按照渠道类型汇总 成功返回结果
type AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_summary_sale_qty_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse
func GetAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse() *AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse {
	return poolAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse.Get().(*AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse 将 AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse(v *AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSummarySaleQtyGetAPIResponse.Put(v)
}
