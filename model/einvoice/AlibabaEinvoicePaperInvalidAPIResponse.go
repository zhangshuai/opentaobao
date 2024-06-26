package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePaperInvalidAPIResponse 纸票作废接口 API返回值
// alibaba.einvoice.paper.invalid
//
// 作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票
type AlibabaEinvoicePaperInvalidAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePaperInvalidAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperInvalidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePaperInvalidAPIResponseModel).Reset()
}

// AlibabaEinvoicePaperInvalidAPIResponseModel is 纸票作废接口 成功返回结果
type AlibabaEinvoicePaperInvalidAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_paper_invalid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用是否成功，操作结果tmc异步返回alibaba_invoice_PaperOpsReturn
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePaperInvalidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoicePaperInvalidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePaperInvalidAPIResponse)
	},
}

// GetAlibabaEinvoicePaperInvalidAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePaperInvalidAPIResponse
func GetAlibabaEinvoicePaperInvalidAPIResponse() *AlibabaEinvoicePaperInvalidAPIResponse {
	return poolAlibabaEinvoicePaperInvalidAPIResponse.Get().(*AlibabaEinvoicePaperInvalidAPIResponse)
}

// ReleaseAlibabaEinvoicePaperInvalidAPIResponse 将 AlibabaEinvoicePaperInvalidAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePaperInvalidAPIResponse(v *AlibabaEinvoicePaperInvalidAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePaperInvalidAPIResponse.Put(v)
}
