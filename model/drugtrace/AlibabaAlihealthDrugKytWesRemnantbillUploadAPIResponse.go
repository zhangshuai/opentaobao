package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytwesremnantbilluploadAPIResponse wes零头出入库单据上传 API返回值
// alibaba.alihealth.drug.kyt.wes.remnantbill.upload
//
// wes零头出入库单据上传
type AlibabaalihealthdrugkytwesremnantbilluploadAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytwesremnantbilluploadAPIResponseModel
}

// AlibabaalihealthdrugkytwesremnantbilluploadAPIResponseModel is wes零头出入库单据上传 成功返回结果
type AlibabaalihealthdrugkytwesremnantbilluploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_wes_remnantbill_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseStatus bool `json:"response_status,omitempty" xml:"response_status,omitempty"`
}