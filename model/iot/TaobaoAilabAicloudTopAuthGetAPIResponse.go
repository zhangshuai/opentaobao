package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopauthgetAPIResponse 登陆 API返回值
// taobao.ailab.aicloud.top.auth.get
//
// 登陆
type TaobaoailabaicloudtopauthgetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopauthgetAPIResponseModel
}

// TaobaoailabaicloudtopauthgetAPIResponseModel is 登陆 成功返回结果
type TaobaoailabaicloudtopauthgetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_auth_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
