package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest 德比付费会员卡开票自匹配 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.auto
//
// 德比付费会员卡开票自匹配
type AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 关键词
	_keyword string
}

// NewAlitripmerchantgalaxyderbymembervoucherreceiptautoRequest 初始化AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherreceiptautoRequest() *AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.auto"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetKeyword is Keyword Setter
// 关键词
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest) GetKeyword() string {
	return r._keyword
}