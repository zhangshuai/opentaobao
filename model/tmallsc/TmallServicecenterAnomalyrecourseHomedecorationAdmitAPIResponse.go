package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponse 天猫服务平台商家投诉单服务商认责接口 API返回值
// tmall.servicecenter.anomalyrecourse.homedecoration.admit
//
// 天猫服务平台商家投诉单服务商认责接口
type TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponse struct {
	model.CommonResponse
	TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponseModel
}

// TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponseModel is 天猫服务平台商家投诉单服务商认责接口 成功返回结果
type TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_anomalyrecourse_homedecoration_admit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallservicecenteranomalyrecoursehomedecorationadmitResult `json:"result,omitempty" xml:"result,omitempty"`
}