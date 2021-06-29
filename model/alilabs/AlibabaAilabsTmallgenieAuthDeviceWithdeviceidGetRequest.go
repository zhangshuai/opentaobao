package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据三方ID查询设备注册激活信息 API请求
alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get

根据三方ID查询设备注册激活信息
*/
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest struct {
    model.Params
    // 设备产品ID
    clientId   string
    // mac地址
    mac   string
}

// 初始化AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest{
    return &AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ClientId Setter
// 设备产品ID
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) SetClientId(clientId string) error {
    r.clientId = clientId
    r.Set("client_id", clientId)
    return nil
}

// ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetClientId() string {
    return r.clientId
}
// Mac Setter
// mac地址
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest) GetMac() string {
    return r.mac
}
