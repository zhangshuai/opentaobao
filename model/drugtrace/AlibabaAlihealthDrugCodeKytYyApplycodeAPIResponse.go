package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytyyapplycodeAPIResponse 医院药品子码申请接口 API返回值
// alibaba.alihealth.drug.code.kyt.yy.applycode
//
// 根据父码及所属企业ID生成子码信息
type AlibabaalihealthdrugcodekytyyapplycodeAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugcodekytyyapplycodeAPIResponseModel
}

// AlibabaalihealthdrugcodekytyyapplycodeAPIResponseModel is 医院药品子码申请接口 成功返回结果
type AlibabaalihealthdrugcodekytyyapplycodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_yy_applycode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaalihealthdrugcodekytyyapplycodeResult `json:"result,omitempty" xml:"result,omitempty"`
}
