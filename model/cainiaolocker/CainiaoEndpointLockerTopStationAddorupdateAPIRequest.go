package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加更新代收点 API请求
cainiao.endpoint.locker.top.station.addorupdate

新增或者修改代收点相关信息
*/
type CainiaoEndpointLockerTopStationAddorupdateAPIRequest struct {
    model.Params
    // 站点信息
    _stationInfo   *StationInfo
}

// 初始化CainiaoEndpointLockerTopStationAddorupdateAPIRequest对象
func NewCainiaoEndpointLockerTopStationAddorupdateRequest() *CainiaoEndpointLockerTopStationAddorupdateAPIRequest{
    return &CainiaoEndpointLockerTopStationAddorupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.station.addorupdate"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StationInfo Setter
// 站点信息
func (r *CainiaoEndpointLockerTopStationAddorupdateAPIRequest) SetStationInfo(_stationInfo *StationInfo) error {
    r._stationInfo = _stationInfo
    r.Set("station_info", _stationInfo)
    return nil
}

// StationInfo Getter
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetStationInfo() *StationInfo {
    return r._stationInfo
}