package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymemberloginderbyAPIRequest 小程序通过德比登入（会员认证） API请求
// alitrip.merchant.galaxy.member.login.derby
//
// 会员认证(德比切换接口，包含授权接口)   返回认证及授权状态
type AlitripmerchantgalaxymemberloginderbyAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// token
	_token string
	// 场景来源，记录什么场景调用什么接口
	_sceneSource string
	// 德比认证会员信息入参
	_derbyAuthenticationParam *DerbyAuthenticationParam
}

// NewAlitripmerchantgalaxymemberloginderbyRequest 初始化AlitripmerchantgalaxymemberloginderbyAPIRequest对象
func NewAlitripmerchantgalaxymemberloginderbyRequest() *AlitripmerchantgalaxymemberloginderbyAPIRequest {
	return &AlitripmerchantgalaxymemberloginderbyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.member.login.derby"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxymemberloginderbyAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// token
func (r *AlitripmerchantgalaxymemberloginderbyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetToken() string {
	return r._token
}

// SetSceneSource is SceneSource Setter
// 场景来源，记录什么场景调用什么接口
func (r *AlitripmerchantgalaxymemberloginderbyAPIRequest) SetSceneSource(_sceneSource string) error {
	r._sceneSource = _sceneSource
	r.Set("scene_source", _sceneSource)
	return nil
}

// GetSceneSource SceneSource Getter
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetSceneSource() string {
	return r._sceneSource
}

// SetDerbyAuthenticationParam is DerbyAuthenticationParam Setter
// 德比认证会员信息入参
func (r *AlitripmerchantgalaxymemberloginderbyAPIRequest) SetDerbyAuthenticationParam(_derbyAuthenticationParam *DerbyAuthenticationParam) error {
	r._derbyAuthenticationParam = _derbyAuthenticationParam
	r.Set("derby_authentication_param", _derbyAuthenticationParam)
	return nil
}

// GetDerbyAuthenticationParam DerbyAuthenticationParam Getter
func (r AlitripmerchantgalaxymemberloginderbyAPIRequest) GetDerbyAuthenticationParam() *DerbyAuthenticationParam {
	return r._derbyAuthenticationParam
}
