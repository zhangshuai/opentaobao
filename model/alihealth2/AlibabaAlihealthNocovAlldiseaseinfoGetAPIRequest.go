package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全国疫情统计数据 API请求
alibaba.alihealth.nocov.alldiseaseinfo.get

获取全国疫情统计数据
*/
type AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest struct {
    model.Params
    // 省的
    _province   string
    // 城市
    _city   string
    // 城市code
    _cityCode   string
}

// 初始化AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest对象
func NewAlibabaAlihealthNocovAlldiseaseinfoGetRequest() *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest{
    return &AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nocov.alldiseaseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Province Setter
// 省的
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetProvince() string {
    return r._province
}
// City Setter
// 城市
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetCity() string {
    return r._city
}
// CityCode Setter
// 城市code
func (r *AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthNocovAlldiseaseinfoGetAPIRequest) GetCityCode() string {
    return r._cityCode
}