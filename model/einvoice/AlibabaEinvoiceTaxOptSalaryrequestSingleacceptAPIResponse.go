package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse 单明细发薪受理 API返回值
// alibaba.einvoice.tax.opt.salaryrequest.singleaccept
//
// 单明细发薪受理
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponseModel).Reset()
}

// AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponseModel is 单明细发薪受理 成功返回结果
type AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryrequest_singleaccept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果类
	Result *TaxOptimizationSingleDetailPaymentAccessResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse)
	},
}

// GetAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse
func GetAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse() *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse {
	return poolAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse.Get().(*AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse)
}

// ReleaseAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse 将 AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse(v *AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse.Put(v)
}
