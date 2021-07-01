package baodian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取宝点SDK的配置项（已迁移） API请求
taobao.baodian.server.sdk.config.get

获取SDK各种配置项（已迁移）
*/
type TaobaoBaodianServerSdkConfigGetAPIRequest struct {
    model.Params
    // appKey
    _appkey   string
    // 渠道
    _channel   string
    // sdk版本号
    _sdkVer   string
    // 与后端约定
    _type   int64
}

// 初始化TaobaoBaodianServerSdkConfigGetAPIRequest对象
func NewTaobaoBaodianServerSdkConfigGetRequest() *TaobaoBaodianServerSdkConfigGetAPIRequest{
    return &TaobaoBaodianServerSdkConfigGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaodianServerSdkConfigGetAPIRequest) GetApiMethodName() string {
    return "taobao.baodian.server.sdk.config.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaodianServerSdkConfigGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Appkey Setter
// appKey
func (r *TaobaoBaodianServerSdkConfigGetAPIRequest) SetAppkey(_appkey string) error {
    r._appkey = _appkey
    r.Set("appkey", _appkey)
    return nil
}

// Appkey Getter
func (r TaobaoBaodianServerSdkConfigGetAPIRequest) GetAppkey() string {
    return r._appkey
}
// Channel Setter
// 渠道
func (r *TaobaoBaodianServerSdkConfigGetAPIRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoBaodianServerSdkConfigGetAPIRequest) GetChannel() string {
    return r._channel
}
// SdkVer Setter
// sdk版本号
func (r *TaobaoBaodianServerSdkConfigGetAPIRequest) SetSdkVer(_sdkVer string) error {
    r._sdkVer = _sdkVer
    r.Set("sdk_ver", _sdkVer)
    return nil
}

// SdkVer Getter
func (r TaobaoBaodianServerSdkConfigGetAPIRequest) GetSdkVer() string {
    return r._sdkVer
}
// Type Setter
// 与后端约定
func (r *TaobaoBaodianServerSdkConfigGetAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoBaodianServerSdkConfigGetAPIRequest) GetType() int64 {
    return r._type
}