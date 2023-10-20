package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest v5.0付费会员卡开票抬头自匹配 API请求
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.auto
//
// v5.0付费会员卡开票抬头自匹配
type AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest struct {
	model.Params
	// tenant_key
	_tenantKey string
	// 关键词
	_keyword string
}

// NewAlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoRequest 初始化AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest对象
func NewAlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoRequest() *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest {
	return &AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.derby.member.voucher.receipt.details.auto"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// tenant_key
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetKeyword is Keyword Setter
// 关键词
func (r *AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest) GetKeyword() string {
	return r._keyword
}