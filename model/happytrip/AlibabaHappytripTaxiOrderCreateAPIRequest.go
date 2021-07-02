package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderCreateAPIRequest 用户叫车 API请求
// alibaba.happytrip.taxi.order.create
//
// 用户根据需要发起叫车请求，在发起请求之前必须事先获得order id.
type AlibabaHappytripTaxiOrderCreateAPIRequest struct {
	model.Params
	// 用户唯一标识
	_uid string
	// 请求id 获取请参见
	_orderId string
	// 叫车车型，0(实时)；1(预约)
	_type int64
	// 乘客手机号
	_passengerPhone string
	// 出发地城市
	_city string
	// 出发地纬度
	_flat string
	// 出发地经度
	_flng string
	// 出发地名称(最多50个字)
	_startName string
	// 出发地详细地址(最多100个字)
	_startAddress string
	// 目的地纬度
	_tlat string
	// 目的地经度
	_tlng string
	// 目的地名称(最多50个字)
	_endName string
	// 目的地详细地址(最多100个字)
	_endAddress string
	// 当前位置纬度
	_clat string
	// 当前位置经度
	_clng string
	// 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
	_departureTime string
	// 车型代码
	_requireLevel string
	// 客户端时间（例如：2015-06-16 12:00:09）
	_appTime string
	// 地图类型:amap：高德，默认高德地图
	_mapType string
	// 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
	_smsPolicy int64
	// 备注
	_extraInfo string
	// 价格md5,通过 预估价接口获得
	_dynamicMd5 string
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
	// 线路类型，0或空表示普通线路；1，表示一口价路线
	_lineType int64
	// 0：不拼车 1:允许拼车，默认不拼车
	_carpoolType int64
	// 乘车人数，允许拼车时该参数不为空
	_passengerNumber int64
}

// NewAlibabaHappytripTaxiOrderCreateRequest 初始化AlibabaHappytripTaxiOrderCreateAPIRequest对象
func NewAlibabaHappytripTaxiOrderCreateRequest() *AlibabaHappytripTaxiOrderCreateAPIRequest {
	return &AlibabaHappytripTaxiOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// Get Uid Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetUid() string {
	return r._uid
}

// Set is OrderId Setter
// 请求id 获取请参见
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is Type Setter
// 叫车车型，0(实时)；1(预约)
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetType() int64 {
	return r._type
}

// Set is PassengerPhone Setter
// 乘客手机号
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetPassengerPhone(_passengerPhone string) error {
	r._passengerPhone = _passengerPhone
	r.Set("passenger_phone", _passengerPhone)
	return nil
}

// Get PassengerPhone Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetPassengerPhone() string {
	return r._passengerPhone
}

// Set is City Setter
// 出发地城市
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// Get City Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetCity() string {
	return r._city
}

// Set is Flat Setter
// 出发地纬度
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetFlat(_flat string) error {
	r._flat = _flat
	r.Set("flat", _flat)
	return nil
}

// Get Flat Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetFlat() string {
	return r._flat
}

// Set is Flng Setter
// 出发地经度
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetFlng(_flng string) error {
	r._flng = _flng
	r.Set("flng", _flng)
	return nil
}

// Get Flng Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetFlng() string {
	return r._flng
}

// Set is StartName Setter
// 出发地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetStartName(_startName string) error {
	r._startName = _startName
	r.Set("start_name", _startName)
	return nil
}

// Get StartName Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetStartName() string {
	return r._startName
}

// Set is StartAddress Setter
// 出发地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetStartAddress(_startAddress string) error {
	r._startAddress = _startAddress
	r.Set("start_address", _startAddress)
	return nil
}

// Get StartAddress Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetStartAddress() string {
	return r._startAddress
}

// Set is Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetTlat(_tlat string) error {
	r._tlat = _tlat
	r.Set("tlat", _tlat)
	return nil
}

// Get Tlat Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetTlat() string {
	return r._tlat
}

// Set is Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetTlng(_tlng string) error {
	r._tlng = _tlng
	r.Set("tlng", _tlng)
	return nil
}

// Get Tlng Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetTlng() string {
	return r._tlng
}

// Set is EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetEndName(_endName string) error {
	r._endName = _endName
	r.Set("end_name", _endName)
	return nil
}

// Get EndName Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetEndName() string {
	return r._endName
}

// Set is EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetEndAddress(_endAddress string) error {
	r._endAddress = _endAddress
	r.Set("end_address", _endAddress)
	return nil
}

// Get EndAddress Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetEndAddress() string {
	return r._endAddress
}

// Set is Clat Setter
// 当前位置纬度
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetClat(_clat string) error {
	r._clat = _clat
	r.Set("clat", _clat)
	return nil
}

// Get Clat Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetClat() string {
	return r._clat
}

// Set is Clng Setter
// 当前位置经度
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetClng(_clng string) error {
	r._clng = _clng
	r.Set("clng", _clng)
	return nil
}

// Get Clng Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetClng() string {
	return r._clng
}

// Set is DepartureTime Setter
// 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetDepartureTime(_departureTime string) error {
	r._departureTime = _departureTime
	r.Set("departure_time", _departureTime)
	return nil
}

// Get DepartureTime Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetDepartureTime() string {
	return r._departureTime
}

// Set is RequireLevel Setter
// 车型代码
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetRequireLevel(_requireLevel string) error {
	r._requireLevel = _requireLevel
	r.Set("require_level", _requireLevel)
	return nil
}

// Get RequireLevel Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetRequireLevel() string {
	return r._requireLevel
}

// Set is AppTime Setter
// 客户端时间（例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetAppTime(_appTime string) error {
	r._appTime = _appTime
	r.Set("app_time", _appTime)
	return nil
}

// Get AppTime Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetAppTime() string {
	return r._appTime
}

// Set is MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetMapType(_mapType string) error {
	r._mapType = _mapType
	r.Set("map_type", _mapType)
	return nil
}

// Get MapType Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetMapType() string {
	return r._mapType
}

// Set is SmsPolicy Setter
// 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetSmsPolicy(_smsPolicy int64) error {
	r._smsPolicy = _smsPolicy
	r.Set("sms_policy", _smsPolicy)
	return nil
}

// Get SmsPolicy Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetSmsPolicy() int64 {
	return r._smsPolicy
}

// Set is ExtraInfo Setter
// 备注
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetExtraInfo(_extraInfo string) error {
	r._extraInfo = _extraInfo
	r.Set("extra_info", _extraInfo)
	return nil
}

// Get ExtraInfo Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetExtraInfo() string {
	return r._extraInfo
}

// Set is DynamicMd5 Setter
// 价格md5,通过 预估价接口获得
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetDynamicMd5(_dynamicMd5 string) error {
	r._dynamicMd5 = _dynamicMd5
	r.Set("dynamic_md5", _dynamicMd5)
	return nil
}

// Get DynamicMd5 Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetDynamicMd5() string {
	return r._dynamicMd5
}

// Set is CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetCostCenter(_costCenter string) error {
	r._costCenter = _costCenter
	r.Set("cost_center", _costCenter)
	return nil
}

// Get CostCenter Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetCostCenter() string {
	return r._costCenter
}

// Set is LineType Setter
// 线路类型，0或空表示普通线路；1，表示一口价路线
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetLineType(_lineType int64) error {
	r._lineType = _lineType
	r.Set("line_type", _lineType)
	return nil
}

// Get LineType Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetLineType() int64 {
	return r._lineType
}

// Set is CarpoolType Setter
// 0：不拼车 1:允许拼车，默认不拼车
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetCarpoolType(_carpoolType int64) error {
	r._carpoolType = _carpoolType
	r.Set("carpool_type", _carpoolType)
	return nil
}

// Get CarpoolType Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetCarpoolType() int64 {
	return r._carpoolType
}

// Set is PassengerNumber Setter
// 乘车人数，允许拼车时该参数不为空
func (r *AlibabaHappytripTaxiOrderCreateAPIRequest) SetPassengerNumber(_passengerNumber int64) error {
	r._passengerNumber = _passengerNumber
	r.Set("passenger_number", _passengerNumber)
	return nil
}

// Get PassengerNumber Getter
func (r AlibabaHappytripTaxiOrderCreateAPIRequest) GetPassengerNumber() int64 {
	return r._passengerNumber
}
