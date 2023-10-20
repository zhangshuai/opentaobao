package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyorderbookAPIRequest 星河-订单预订接口 API请求
// alitrip.merchant.galaxy.order.book
//
// 为星河酒店解决方案的C端用户提供酒店预订能力
type AlitripmerchantgalaxyorderbookAPIRequest struct {
	model.Params
	// 租户身份信息
	_tenantKey string
	// 用户登录token
	_token string
	// 订单编号
	_orderCode string
	// 广告追踪信息
	_sourceQuery string
	// 预订参数
	_orderParam *CreateOrderParam
}

// NewAlitripmerchantgalaxyorderbookRequest 初始化AlitripmerchantgalaxyorderbookAPIRequest对象
func NewAlitripmerchantgalaxyorderbookRequest() *AlitripmerchantgalaxyorderbookAPIRequest {
	return &AlitripmerchantgalaxyorderbookAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.order.book"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户身份信息
func (r *AlitripmerchantgalaxyorderbookAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetToken is Token Setter
// 用户登录token
func (r *AlitripmerchantgalaxyorderbookAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetToken() string {
	return r._token
}

// SetOrderCode is OrderCode Setter
// 订单编号
func (r *AlitripmerchantgalaxyorderbookAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetSourceQuery is SourceQuery Setter
// 广告追踪信息
func (r *AlitripmerchantgalaxyorderbookAPIRequest) SetSourceQuery(_sourceQuery string) error {
	r._sourceQuery = _sourceQuery
	r.Set("source_query", _sourceQuery)
	return nil
}

// GetSourceQuery SourceQuery Getter
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetSourceQuery() string {
	return r._sourceQuery
}

// SetOrderParam is OrderParam Setter
// 预订参数
func (r *AlitripmerchantgalaxyorderbookAPIRequest) SetOrderParam(_orderParam *CreateOrderParam) error {
	r._orderParam = _orderParam
	r.Set("order_param", _orderParam)
	return nil
}

// GetOrderParam OrderParam Getter
func (r AlitripmerchantgalaxyorderbookAPIRequest) GetOrderParam() *CreateOrderParam {
	return r._orderParam
}
