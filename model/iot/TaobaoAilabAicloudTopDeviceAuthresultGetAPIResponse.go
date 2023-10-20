package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceauthresultgetAPIResponse 获取设备授权码验证结果 API返回值
// taobao.ailab.aicloud.top.device.authresult.get
//
// 获取设备授权码验证结果
type TaobaoailabaicloudtopdeviceauthresultgetAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopdeviceauthresultgetAPIResponseModel
}

// TaobaoailabaicloudtopdeviceauthresultgetAPIResponseModel is 获取设备授权码验证结果 成功返回结果
type TaobaoailabaicloudtopdeviceauthresultgetAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_device_authresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
