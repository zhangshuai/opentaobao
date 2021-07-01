package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统帐号登录注册token验证接口 API请求
alibaba.alisports.passport.account.tokenvalidate

阿里体育会员系统帐号登录注册token验证接口
*/
type AlibabaAlisportsPassportAccountTokenvalidateAPIRequest struct {
    model.Params
    // 业务方appkey
    _alispAppKey   string
    // 签名
    _alispSign   string
    // token
    _token   string
    // 注册用户类型
    _userType   int64
    // 时间戳
    _alispTime   string
    // 一键登录参数
    _secret   string
    // json字符串，传入扩展字段
    _extInfo   string
    // 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
    _mtopAppkey   string
}

// 初始化AlibabaAlisportsPassportAccountTokenvalidateAPIRequest对象
func NewAlibabaAlisportsPassportAccountTokenvalidateRequest() *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest{
    return &AlibabaAlisportsPassportAccountTokenvalidateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.tokenvalidate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 业务方appkey
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispSign Setter
// 签名
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetAlispSign() string {
    return r._alispSign
}
// Token Setter
// token
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetToken() string {
    return r._token
}
// UserType Setter
// 注册用户类型
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetUserType(_userType int64) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetUserType() int64 {
    return r._userType
}
// AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetAlispTime() string {
    return r._alispTime
}
// Secret Setter
// 一键登录参数
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetSecret(_secret string) error {
    r._secret = _secret
    r.Set("secret", _secret)
    return nil
}

// Secret Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetSecret() string {
    return r._secret
}
// ExtInfo Setter
// json字符串，传入扩展字段
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetExtInfo(_extInfo string) error {
    r._extInfo = _extInfo
    r.Set("ext_info", _extInfo)
    return nil
}

// ExtInfo Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetExtInfo() string {
    return r._extInfo
}
// MtopAppkey Setter
// 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
func (r *AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) SetMtopAppkey(_mtopAppkey string) error {
    r._mtopAppkey = _mtopAppkey
    r.Set("mtop_appkey", _mtopAppkey)
    return nil
}

// MtopAppkey Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateAPIRequest) GetMtopAppkey() string {
    return r._mtopAppkey
}