package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-校验token API请求
alitrip.merchant.galaxy.member.token

校验或者刷新token
*/
type AlitripMerchantGalaxyMemberTokenAPIRequest struct {
    model.Params
    // 租户身份信息
    _tenantKey   string
    // 访问携带toekn
    _token   string
}

// 初始化AlitripMerchantGalaxyMemberTokenAPIRequest对象
func NewAlitripMerchantGalaxyMemberTokenRequest() *AlitripMerchantGalaxyMemberTokenAPIRequest{
    return &AlitripMerchantGalaxyMemberTokenAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.member.token"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyMemberTokenAPIRequest) SetTenantKey(_tenantKey string) error {
    r._tenantKey = _tenantKey
    r.Set("tenant_key", _tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetTenantKey() string {
    return r._tenantKey
}
// Token Setter
// 访问携带toekn
func (r *AlitripMerchantGalaxyMemberTokenAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyMemberTokenAPIRequest) GetToken() string {
    return r._token
}