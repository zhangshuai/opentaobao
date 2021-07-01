package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小件员GPS位置信息 API请求
cainiao.guoguo.cp.nborderfrontr.uploadcoordinate

上传小件员GPS位置信息
*/
type CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest struct {
    model.Params
    // 小件员所在公司编号
    _cpCode   string
    // 小件员员工编号
    _cpUserId   string
    // 上报时间，格式：yyyy-MM-dd HH:mm:ss
    _timeStamp   string
    // 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
    _source   int64
    // 经度
    _lng   string
    // 纬度
    _lat   string
    // 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
    _gpsType   string
}

// 初始化CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest对象
func NewCainiaoGuoguoCpNborderfrontrUploadcoordinateRequest() *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest{
    return &CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetApiMethodName() string {
    return "cainiao.guoguo.cp.nborderfrontr.uploadcoordinate"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 小件员所在公司编号
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetCpCode() string {
    return r._cpCode
}
// CpUserId Setter
// 小件员员工编号
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetCpUserId(_cpUserId string) error {
    r._cpUserId = _cpUserId
    r.Set("cp_user_id", _cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetCpUserId() string {
    return r._cpUserId
}
// TimeStamp Setter
// 上报时间，格式：yyyy-MM-dd HH:mm:ss
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetTimeStamp(_timeStamp string) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetTimeStamp() string {
    return r._timeStamp
}
// Source Setter
// 来源：1.小件员app sdk 2.驿站 3. 裹裹 10001.圆通行者
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetSource() int64 {
    return r._source
}
// Lng Setter
// 经度
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetLng(_lng string) error {
    r._lng = _lng
    r.Set("lng", _lng)
    return nil
}

// Lng Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetLng() string {
    return r._lng
}
// Lat Setter
// 纬度
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetLat(_lat string) error {
    r._lat = _lat
    r.Set("lat", _lat)
    return nil
}

// Lat Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetLat() string {
    return r._lat
}
// GpsType Setter
// 0 安卓定位，     1 苹果定位，  2 其他系统定位，   10 高德定位，  11 百度定位，  12 google定位     13 其他
func (r *CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) SetGpsType(_gpsType string) error {
    r._gpsType = _gpsType
    r.Set("gps_type", _gpsType)
    return nil
}

// GpsType Getter
func (r CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIRequest) GetGpsType() string {
    return r._gpsType
}