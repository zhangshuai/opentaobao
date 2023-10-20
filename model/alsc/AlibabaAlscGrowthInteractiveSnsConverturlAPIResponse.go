package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscgrowthinteractivesnsconverturlAPIResponse 防封接口 API返回值
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
type AlibabaalscgrowthinteractivesnsconverturlAPIResponse struct {
	model.CommonResponse
	AlibabaalscgrowthinteractivesnsconverturlAPIResponseModel
}

// AlibabaalscgrowthinteractivesnsconverturlAPIResponseModel is 防封接口 成功返回结果
type AlibabaalscgrowthinteractivesnsconverturlAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_sns_converturl_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应
	Result *BaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}