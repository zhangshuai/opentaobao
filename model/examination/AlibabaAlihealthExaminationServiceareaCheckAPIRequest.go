package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationserviceareacheckAPIRequest 体检机构对接_上门检测服务范围查询 API请求
// alibaba.alihealth.examination.servicearea.check
//
// 体检机构对接_上门检测服务范围查询
type AlibabaalihealthexaminationserviceareacheckAPIRequest struct {
	model.Params
	// 机构套餐编码
	_packageCode string
	// 上门检测地址
	_address string
	// 上门检测地址纬度
	_latitude string
	// 上门检测地址经度
	_longitude string
	// 省份名称（高德）
	_province string
	// 省份编码（高德adcode）
	_provinceCode string
	// 城市名称（高德）
	_city string
	// 城市编码（高德adcode）
	_cityCode string
	// 区域名称（高德）
	_district string
	// 区域编码（高德adcode）
	_districtCode string
}

// NewAlibabaalihealthexaminationserviceareacheckRequest 初始化AlibabaalihealthexaminationserviceareacheckAPIRequest对象
func NewAlibabaalihealthexaminationserviceareacheckRequest() *AlibabaalihealthexaminationserviceareacheckAPIRequest {
	return &AlibabaalihealthexaminationserviceareacheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.servicearea.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageCode is PackageCode Setter
// 机构套餐编码
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetPackageCode(_packageCode string) error {
	r._packageCode = _packageCode
	r.Set("package_code", _packageCode)
	return nil
}

// GetPackageCode PackageCode Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetPackageCode() string {
	return r._packageCode
}

// SetAddress is Address Setter
// 上门检测地址
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetAddress() string {
	return r._address
}

// SetLatitude is Latitude Setter
// 上门检测地址纬度
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 上门检测地址经度
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetProvince is Province Setter
// 省份名称（高德）
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetProvince(_province string) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetProvince() string {
	return r._province
}

// SetProvinceCode is ProvinceCode Setter
// 省份编码（高德adcode）
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetProvinceCode(_provinceCode string) error {
	r._provinceCode = _provinceCode
	r.Set("province_code", _provinceCode)
	return nil
}

// GetProvinceCode ProvinceCode Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetProvinceCode() string {
	return r._provinceCode
}

// SetCity is City Setter
// 城市名称（高德）
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetCity() string {
	return r._city
}

// SetCityCode is CityCode Setter
// 城市编码（高德adcode）
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetDistrict is District Setter
// 区域名称（高德）
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetDistrict(_district string) error {
	r._district = _district
	r.Set("district", _district)
	return nil
}

// GetDistrict District Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetDistrict() string {
	return r._district
}

// SetDistrictCode is DistrictCode Setter
// 区域编码（高德adcode）
func (r *AlibabaalihealthexaminationserviceareacheckAPIRequest) SetDistrictCode(_districtCode string) error {
	r._districtCode = _districtCode
	r.Set("district_code", _districtCode)
	return nil
}

// GetDistrictCode DistrictCode Getter
func (r AlibabaalihealthexaminationserviceareacheckAPIRequest) GetDistrictCode() string {
	return r._districtCode
}
