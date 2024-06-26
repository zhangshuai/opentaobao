package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionOrderGetAPIResponse get order list API返回值
// aliexpress.solution.order.get
//
// Get Order List from AliExpress
type AliexpressSolutionOrderGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionOrderGetAPIResponseModel).Reset()
}

// AliexpressSolutionOrderGetAPIResponseModel is get order list 成功返回结果
type AliexpressSolutionOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PaginationResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionOrderGetAPIResponse)
	},
}

// GetAliexpressSolutionOrderGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionOrderGetAPIResponse
func GetAliexpressSolutionOrderGetAPIResponse() *AliexpressSolutionOrderGetAPIResponse {
	return poolAliexpressSolutionOrderGetAPIResponse.Get().(*AliexpressSolutionOrderGetAPIResponse)
}

// ReleaseAliexpressSolutionOrderGetAPIResponse 将 AliexpressSolutionOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionOrderGetAPIResponse(v *AliexpressSolutionOrderGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionOrderGetAPIResponse.Put(v)
}
