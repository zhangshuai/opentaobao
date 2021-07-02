package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerGetAPIResponse 查询顾客详情 API返回值
// alibaba.alsc.crm.customer.get
//
// 查询顾客详情
type AlibabaAlscCrmCustomerGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCustomerGetAPIResponseModel
}

// AlibabaAlscCrmCustomerGetAPIResponseModel is 查询顾客详情 成功返回结果
type AlibabaAlscCrmCustomerGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_customer_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
