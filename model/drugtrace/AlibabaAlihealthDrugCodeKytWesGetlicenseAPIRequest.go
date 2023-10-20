package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest 获取licenseToken API请求
// alibaba.alihealth.drug.code.kyt.wes.getlicense
//
// 获取licenseToken
type AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// license
	_license string
}

// NewAlibabaalihealthdrugcodekytwesgetlicenseRequest 初始化AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest对象
func NewAlibabaalihealthdrugcodekytwesgetlicenseRequest() *AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest {
	return &AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.wes.getlicense"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicense is License Setter
// license
func (r *AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) SetLicense(_license string) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r AlibabaalihealthdrugcodekytwesgetlicenseAPIRequest) GetLicense() string {
	return r._license
}
