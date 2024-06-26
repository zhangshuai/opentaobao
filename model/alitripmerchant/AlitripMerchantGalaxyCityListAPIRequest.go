package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCityListAPIRequest 星河-酒店城市列表展示 API请求
// alitrip.merchant.galaxy.city.list
//
// 雅高酒店城市列表展示，并且首字母列出酒店城市
type AlitripMerchantGalaxyCityListAPIRequest struct {
	model.Params
	// 商家租户id
	_tenantKey string
	// 0国内 1国外
	_domestic int64
}

// NewAlitripMerchantGalaxyCityListRequest 初始化AlitripMerchantGalaxyCityListAPIRequest对象
func NewAlitripMerchantGalaxyCityListRequest() *AlitripMerchantGalaxyCityListAPIRequest {
	return &AlitripMerchantGalaxyCityListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyCityListAPIRequest) Reset() {
	r._tenantKey = ""
	r._domestic = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyCityListAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.city.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyCityListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyCityListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 商家租户id
func (r *AlitripMerchantGalaxyCityListAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyCityListAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetDomestic is Domestic Setter
// 0国内 1国外
func (r *AlitripMerchantGalaxyCityListAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// GetDomestic Domestic Getter
func (r AlitripMerchantGalaxyCityListAPIRequest) GetDomestic() int64 {
	return r._domestic
}

var poolAlitripMerchantGalaxyCityListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyCityListRequest()
	},
}

// GetAlitripMerchantGalaxyCityListRequest 从 sync.Pool 获取 AlitripMerchantGalaxyCityListAPIRequest
func GetAlitripMerchantGalaxyCityListAPIRequest() *AlitripMerchantGalaxyCityListAPIRequest {
	return poolAlitripMerchantGalaxyCityListAPIRequest.Get().(*AlitripMerchantGalaxyCityListAPIRequest)
}

// ReleaseAlitripMerchantGalaxyCityListAPIRequest 将 AlitripMerchantGalaxyCityListAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyCityListAPIRequest(v *AlitripMerchantGalaxyCityListAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyCityListAPIRequest.Put(v)
}
