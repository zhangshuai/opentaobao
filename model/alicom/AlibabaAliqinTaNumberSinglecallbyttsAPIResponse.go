package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqintanumbersinglecallbyttsAPIResponse 根据号码tts单呼 API返回值
// alibaba.aliqin.ta.number.singlecallbytts
//
// 将语音验证码和语音通知发布至聚石塔渠道
type AlibabaaliqintanumbersinglecallbyttsAPIResponse struct {
	model.CommonResponse
	AlibabaaliqintanumbersinglecallbyttsAPIResponseModel
}

// AlibabaaliqintanumbersinglecallbyttsAPIResponseModel is 根据号码tts单呼 成功返回结果
type AlibabaaliqintanumbersinglecallbyttsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_ta_number_singlecallbytts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaaliqintanumbersinglecallbyttsResult `json:"result,omitempty" xml:"result,omitempty"`
}
