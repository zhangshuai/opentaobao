package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest 医院发药机 API请求
// alibaba.alihealth.drug.code.kyt.hospitalsenddrugmachine
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
// 提供给医院发药机使用
type AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest struct {
	model.Params
	// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业唯一标识
	_refEntId string
	// 69码
	_barcode69 string
}

// NewAlibabaalihealthdrugcodekythospitalsenddrugmachineRequest 初始化AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest对象
func NewAlibabaalihealthdrugcodekythospitalsenddrugmachineRequest() *AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest {
	return &AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.hospitalsenddrugmachine"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBarcode69 is Barcode69 Setter
// 69码
func (r *AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) SetBarcode69(_barcode69 string) error {
	r._barcode69 = _barcode69
	r.Set("barcode69", _barcode69)
	return nil
}

// GetBarcode69 Barcode69 Getter
func (r AlibabaalihealthdrugcodekythospitalsenddrugmachineAPIRequest) GetBarcode69() string {
	return r._barcode69
}