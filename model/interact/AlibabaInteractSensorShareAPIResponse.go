package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分享 API返回值 
alibaba.interact.sensor.share

客户端分享
*/
type AlibabaInteractSensorShareAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorShareAPIResponseModel
}

// 分享 成功返回结果
type AlibabaInteractSensorShareAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_share_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // return=0表示成功
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}