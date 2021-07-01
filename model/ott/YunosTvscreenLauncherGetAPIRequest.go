package ott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
一体机桌面 API请求
yunos.tvscreen.launcher.get

LCTS一体机桌面后台,提供基于运营坑位适配的桌面服务
*/
type YunosTvscreenLauncherGetAPIRequest struct {
    model.Params
    // 设备属性
    _property   string
    // IP来源
    _ip   string
}

// 初始化YunosTvscreenLauncherGetAPIRequest对象
func NewYunosTvscreenLauncherGetRequest() *YunosTvscreenLauncherGetAPIRequest{
    return &YunosTvscreenLauncherGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvscreenLauncherGetAPIRequest) GetApiMethodName() string {
    return "yunos.tvscreen.launcher.get"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvscreenLauncherGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Property Setter
// 设备属性
func (r *YunosTvscreenLauncherGetAPIRequest) SetProperty(_property string) error {
    r._property = _property
    r.Set("property", _property)
    return nil
}

// Property Getter
func (r YunosTvscreenLauncherGetAPIRequest) GetProperty() string {
    return r._property
}
// Ip Setter
// IP来源
func (r *YunosTvscreenLauncherGetAPIRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r YunosTvscreenLauncherGetAPIRequest) GetIp() string {
    return r._ip
}