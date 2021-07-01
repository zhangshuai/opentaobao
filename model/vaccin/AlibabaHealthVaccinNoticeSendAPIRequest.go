package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送消息提醒 API请求
alibaba.health.vaccin.notice.send

ISV 通过免疫规划中心给用户发送短信或者支付宝 PUSH 提醒。
*/
type AlibabaHealthVaccinNoticeSendAPIRequest struct {
    model.Params
    // 支付宝用户 ID
    _alipayUserId   string
    // ISV 侧用户 ID
    _isvUserId   string
    // 订单 ID
    _orderId   string
    // 接种人姓名
    _name   string
    // 联系电话
    _mobile   string
    // 接种的疫苗信息
    _vaccineList   []VaccineInfo
    // 接种人性别:1=男,2=女
    _sex   int64
    // 接种人出生日期
    _birthday   string
    // 接种点编码
    _povNo   string
    // 接种点名称
    _povName   string
    // 接种点地址
    _address   string
    // 省份名称
    _province   string
    // 城市名称
    _city   string
    // 区名称
    _area   string
    // 预约日期
    _reserveDate   string
    // 预约时间段
    _reserveTime   string
    // 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
    _messageType   int64
    // 用户入口，支付宝或医鹿，alipay或yl
    _appChannel   string
}

// 初始化AlibabaHealthVaccinNoticeSendAPIRequest对象
func NewAlibabaHealthVaccinNoticeSendRequest() *AlibabaHealthVaccinNoticeSendAPIRequest{
    return &AlibabaHealthVaccinNoticeSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.notice.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayUserId Setter
// 支付宝用户 ID
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetAlipayUserId(_alipayUserId string) error {
    r._alipayUserId = _alipayUserId
    r.Set("alipay_user_id", _alipayUserId)
    return nil
}

// AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetAlipayUserId() string {
    return r._alipayUserId
}
// IsvUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetIsvUserId(_isvUserId string) error {
    r._isvUserId = _isvUserId
    r.Set("isv_user_id", _isvUserId)
    return nil
}

// IsvUserId Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetIsvUserId() string {
    return r._isvUserId
}
// OrderId Setter
// 订单 ID
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetOrderId() string {
    return r._orderId
}
// Name Setter
// 接种人姓名
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetName() string {
    return r._name
}
// Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetMobile() string {
    return r._mobile
}
// VaccineList Setter
// 接种的疫苗信息
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetVaccineList(_vaccineList []VaccineInfo) error {
    r._vaccineList = _vaccineList
    r.Set("vaccine_list", _vaccineList)
    return nil
}

// VaccineList Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetVaccineList() []VaccineInfo {
    return r._vaccineList
}
// Sex Setter
// 接种人性别:1=男,2=女
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetSex(_sex int64) error {
    r._sex = _sex
    r.Set("sex", _sex)
    return nil
}

// Sex Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetSex() int64 {
    return r._sex
}
// Birthday Setter
// 接种人出生日期
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetBirthday(_birthday string) error {
    r._birthday = _birthday
    r.Set("birthday", _birthday)
    return nil
}

// Birthday Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetBirthday() string {
    return r._birthday
}
// PovNo Setter
// 接种点编码
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetPovNo(_povNo string) error {
    r._povNo = _povNo
    r.Set("pov_no", _povNo)
    return nil
}

// PovNo Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetPovNo() string {
    return r._povNo
}
// PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetPovName(_povName string) error {
    r._povName = _povName
    r.Set("pov_name", _povName)
    return nil
}

// PovName Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetPovName() string {
    return r._povName
}
// Address Setter
// 接种点地址
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetAddress() string {
    return r._address
}
// Province Setter
// 省份名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetProvince(_province string) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetProvince() string {
    return r._province
}
// City Setter
// 城市名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetCity() string {
    return r._city
}
// Area Setter
// 区名称
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetArea(_area string) error {
    r._area = _area
    r.Set("area", _area)
    return nil
}

// Area Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetArea() string {
    return r._area
}
// ReserveDate Setter
// 预约日期
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetReserveDate() string {
    return r._reserveDate
}
// ReserveTime Setter
// 预约时间段
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetReserveTime(_reserveTime string) error {
    r._reserveTime = _reserveTime
    r.Set("reserve_time", _reserveTime)
    return nil
}

// ReserveTime Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetReserveTime() string {
    return r._reserveTime
}
// MessageType Setter
// 消息提醒类型：1=疫苗下一针预约提醒  2=预约成功提醒  3=接种提醒  4=补种提醒
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetMessageType(_messageType int64) error {
    r._messageType = _messageType
    r.Set("message_type", _messageType)
    return nil
}

// MessageType Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetMessageType() int64 {
    return r._messageType
}
// AppChannel Setter
// 用户入口，支付宝或医鹿，alipay或yl
func (r *AlibabaHealthVaccinNoticeSendAPIRequest) SetAppChannel(_appChannel string) error {
    r._appChannel = _appChannel
    r.Set("app_channel", _appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaHealthVaccinNoticeSendAPIRequest) GetAppChannel() string {
    return r._appChannel
}