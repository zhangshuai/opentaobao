package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccineregistercancelAPIResponse 取消登记 API返回值
// alibaba.alihealth.vaccine.register.cancel
//
// 取消登记
type AlibabaalihealthvaccineregistercancelAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthvaccineregistercancelAPIResponseModel
}

// AlibabaalihealthvaccineregistercancelAPIResponseModel is 取消登记 成功返回结果
type AlibabaalihealthvaccineregistercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_register_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	Result *AlibabaalihealthvaccineregistercancelMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
