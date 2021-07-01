package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 API请求
alibaba.alink.device.unbind

阿里智能解绑设备
*/
type AlibabaAlinkDeviceUnbindAPIRequest struct {
    model.Params
    // 设备id
    _uuid   string
}

// 初始化AlibabaAlinkDeviceUnbindAPIRequest对象
func NewAlibabaAlinkDeviceUnbindRequest() *AlibabaAlinkDeviceUnbindAPIRequest{
    return &AlibabaAlinkDeviceUnbindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetApiMethodName() string {
    return "alibaba.alink.device.unbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceUnbindAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceUnbindAPIRequest) GetUuid() string {
    return r._uuid
}