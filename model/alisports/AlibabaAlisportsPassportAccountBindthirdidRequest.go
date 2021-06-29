package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育三方ID绑定接口 API请求
alibaba.alisports.passport.account.bindthirdid

阿里体育三方ID绑定接口
*/
type AlibabaAlisportsPassportAccountBindthirdidRequest struct {
    model.Params
    // 业务方appkey
    alispAppKey   string
    // 时间戳精确到秒
    alispTime   string
    // 接口签名
    alispSign   string
    // 阿里体育用户ID
    aliuid   string
    // 三方ID
    appUid   string
    // 手机号
    mobile   string
}

// 初始化AlibabaAlisportsPassportAccountBindthirdidRequest对象
func NewAlibabaAlisportsPassportAccountBindthirdidRequest() *AlibabaAlisportsPassportAccountBindthirdidRequest{
    return &AlibabaAlisportsPassportAccountBindthirdidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.bindthirdid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 业务方appkey
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAlispSign() string {
    return r.alispSign
}
// Aliuid Setter
// 阿里体育用户ID
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAliuid() string {
    return r.aliuid
}
// AppUid Setter
// 三方ID
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAppUid(appUid string) error {
    r.appUid = appUid
    r.Set("app_uid", appUid)
    return nil
}

// AppUid Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAppUid() string {
    return r.appUid
}
// Mobile Setter
// 手机号
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetMobile() string {
    return r.mobile
}