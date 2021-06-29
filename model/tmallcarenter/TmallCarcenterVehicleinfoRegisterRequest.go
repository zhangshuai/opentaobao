package tmallcarenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
车型数据更新 API请求
tmall.carcenter.vehicleinfo.register

基本车型信息维护
*/
type TmallCarcenterVehicleinfoRegisterRequest struct {
    model.Params
    // 车型数据对象
    vehicleInfo   *OriginVehicleInfoDto
}

// 初始化TmallCarcenterVehicleinfoRegisterRequest对象
func NewTmallCarcenterVehicleinfoRegisterRequest() *TmallCarcenterVehicleinfoRegisterRequest{
    return &TmallCarcenterVehicleinfoRegisterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarcenterVehicleinfoRegisterRequest) GetApiMethodName() string {
    return "tmall.carcenter.vehicleinfo.register"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarcenterVehicleinfoRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VehicleInfo Setter
// 车型数据对象
func (r *TmallCarcenterVehicleinfoRegisterRequest) SetVehicleInfo(vehicleInfo *OriginVehicleInfoDto) error {
    r.vehicleInfo = vehicleInfo
    r.Set("vehicle_info", vehicleInfo)
    return nil
}

// VehicleInfo Getter
func (r TmallCarcenterVehicleinfoRegisterRequest) GetVehicleInfo() *OriginVehicleInfoDto {
    return r.vehicleInfo
}
