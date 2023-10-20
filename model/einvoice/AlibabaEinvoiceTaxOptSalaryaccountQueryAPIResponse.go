package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponse 查询用户的发薪账号 API返回值
// alibaba.einvoice.tax.opt.salaryaccount.query
//
// 查询用户的发薪账号状态
type AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponseModel
}

// AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponseModel is 查询用户的发薪账号 成功返回结果
type AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salaryaccount_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TaxOptimizationQueryAlipayAccountResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
