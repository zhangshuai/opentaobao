package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开放设备id转换内部设备id API返回值 
taobao.ailab.aicloud.top.device.deviceid.convert

将开放设备id转换为内部设备id
*/
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceDeviceidConvertResponse
}

// 开放设备id转换内部设备id 成功返回结果
type TaobaoAilabAicloudTopDeviceDeviceidConvertResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_deviceid_convert_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceDeviceidConvertResult `json:"result,omitempty" xml:"result,omitempty"`
}
