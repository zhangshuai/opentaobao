package alihealthoutflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowPrescriptionUpdateAPIResponse 处方外流-修改处方 API返回值
// alibaba.alihealth.outflow.prescription.update
//
// 阿里健康-处方外流-对外提供处方修改功能
type AlibabaAlihealthOutflowPrescriptionUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthOutflowPrescriptionUpdateAPIResponseModel
}

// AlibabaAlihealthOutflowPrescriptionUpdateAPIResponseModel is 处方外流-修改处方 成功返回结果
type AlibabaAlihealthOutflowPrescriptionUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_outflow_prescription_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ServiceResult
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
