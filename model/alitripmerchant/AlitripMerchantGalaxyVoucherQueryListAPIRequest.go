package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyvoucherquerylistAPIRequest 查询代金券列表 API请求
// alitrip.merchant.galaxy.voucher.query.list
//
// 查询代金券列表
type AlitripmerchantgalaxyvoucherquerylistAPIRequest struct {
	model.Params
	// 租户id
	_tenantKey string
	// 用户token
	_token string
}

// NewAlitripmerchantgalaxyvoucherquerylistRequest 初始化AlitripmerchantgalaxyvoucherquerylistAPIRequest对象
func NewAlitripmerchantgalaxyvoucherquerylistRequest() *AlitripmerchantgalaxyvoucherquerylistAPIRequest {
	return &AlitripmerchantgalaxyvoucherquerylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyvoucherquerylistAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.voucher.query.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyvoucherquerylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyvoucherquerylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户id
func (r *AlitripmerchantgalaxyvoucherquerylistAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyvoucherquerylistAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户token
func (r *AlitripmerchantgalaxyvoucherquerylistAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyvoucherquerylistAPIRequest) GetToken() string {
	return r._token
}
