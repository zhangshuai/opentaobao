package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取结果接口 API请求
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口
*/
type AlibabaSecurityJaqRpFetchmaterialAPIRequest struct {
    model.Params
    // 消息服务推送的key
    _securityKey   string
}

// 初始化AlibabaSecurityJaqRpFetchmaterialAPIRequest对象
func NewAlibabaSecurityJaqRpFetchmaterialRequest() *AlibabaSecurityJaqRpFetchmaterialAPIRequest{
    return &AlibabaSecurityJaqRpFetchmaterialAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.fetchmaterial"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SecurityKey Setter
// 消息服务推送的key
func (r *AlibabaSecurityJaqRpFetchmaterialAPIRequest) SetSecurityKey(_securityKey string) error {
    r._securityKey = _securityKey
    r.Set("security_key", _securityKey)
    return nil
}

// SecurityKey Getter
func (r AlibabaSecurityJaqRpFetchmaterialAPIRequest) GetSecurityKey() string {
    return r._securityKey
}