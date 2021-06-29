package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵淘宝登录授权绑定接口 API请求
alibaba.ailabs.tmallgenie.auth.taobaoauth

厂商获取用户淘宝授权之后，通过此接口获取天猫精灵授权，并绑定一台设备
*/
type AlibabaAilabsTmallgenieAuthTaobaoauthRequest struct {
    model.Params
    // 授权信息
    authParam   *TopAuthReqDto
    // 设备信息
    deviceParam   *TopDeviceReqDto
}

// 初始化AlibabaAilabsTmallgenieAuthTaobaoauthRequest对象
func NewAlibabaAilabsTmallgenieAuthTaobaoauthRequest() *AlibabaAilabsTmallgenieAuthTaobaoauthRequest{
    return &AlibabaAilabsTmallgenieAuthTaobaoauthRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetApiMethodName() string {
    return "alibaba.ailabs.tmallgenie.auth.taobaoauth"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuthParam Setter
// 授权信息
func (r *AlibabaAilabsTmallgenieAuthTaobaoauthRequest) SetAuthParam(authParam *TopAuthReqDto) error {
    r.authParam = authParam
    r.Set("auth_param", authParam)
    return nil
}

// AuthParam Getter
func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetAuthParam() *TopAuthReqDto {
    return r.authParam
}
// DeviceParam Setter
// 设备信息
func (r *AlibabaAilabsTmallgenieAuthTaobaoauthRequest) SetDeviceParam(deviceParam *TopDeviceReqDto) error {
    r.deviceParam = deviceParam
    r.Set("device_param", deviceParam)
    return nil
}

// DeviceParam Getter
func (r AlibabaAilabsTmallgenieAuthTaobaoauthRequest) GetDeviceParam() *TopDeviceReqDto {
    return r.deviceParam
}