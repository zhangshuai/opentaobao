package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyOrderValidateAPIRequest 星河-订单试单接口 API请求
// alitrip.merchant.galaxy.order.validate
//
// 根据用户选择酒店房型、入住人数、预订时间参数，获取是否可预订及价格变化信息
type AlitripMerchantGalaxyOrderValidateAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户标识
	_token string
	// 试单参数
	_validateOrderParam *ValidateOrderParam
}

// NewAlitripMerchantGalaxyOrderValidateRequest 初始化AlitripMerchantGalaxyOrderValidateAPIRequest对象
func NewAlitripMerchantGalaxyOrderValidateRequest() *AlitripMerchantGalaxyOrderValidateAPIRequest {
	return &AlitripMerchantGalaxyOrderValidateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) Reset() {
	r._tenantKey = ""
	r._token = ""
	r._validateOrderParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户标识
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetToken() string {
	return r._token
}

// SetValidateOrderParam is ValidateOrderParam Setter
// 试单参数
func (r *AlitripMerchantGalaxyOrderValidateAPIRequest) SetValidateOrderParam(_validateOrderParam *ValidateOrderParam) error {
	r._validateOrderParam = _validateOrderParam
	r.Set("validate_order_param", _validateOrderParam)
	return nil
}

// GetValidateOrderParam ValidateOrderParam Getter
func (r AlitripMerchantGalaxyOrderValidateAPIRequest) GetValidateOrderParam() *ValidateOrderParam {
	return r._validateOrderParam
}

var poolAlitripMerchantGalaxyOrderValidateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyOrderValidateRequest()
	},
}

// GetAlitripMerchantGalaxyOrderValidateRequest 从 sync.Pool 获取 AlitripMerchantGalaxyOrderValidateAPIRequest
func GetAlitripMerchantGalaxyOrderValidateAPIRequest() *AlitripMerchantGalaxyOrderValidateAPIRequest {
	return poolAlitripMerchantGalaxyOrderValidateAPIRequest.Get().(*AlitripMerchantGalaxyOrderValidateAPIRequest)
}

// ReleaseAlitripMerchantGalaxyOrderValidateAPIRequest 将 AlitripMerchantGalaxyOrderValidateAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyOrderValidateAPIRequest(v *AlitripMerchantGalaxyOrderValidateAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyOrderValidateAPIRequest.Put(v)
}
