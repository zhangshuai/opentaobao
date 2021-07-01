package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
终端配置支持 API请求
taobao.lark.pos.itemprod.findterminal

终端配置支持,读取如果不存在则创建和远程的连接配置并返回
*/
type TaobaoLarkPosItemprodFindterminalAPIRequest struct {
    model.Params
    // 终端id
    _deviceId   string
    // 终端类型
    _deviceType   string
    // 912874323429834
    _createUser   string
    // 租户编码
    _leaseCode   string
    // 影城id
    _cinemaId   string
    // 影城名称
    _cinemaName   string
}

// 初始化TaobaoLarkPosItemprodFindterminalAPIRequest对象
func NewTaobaoLarkPosItemprodFindterminalRequest() *TaobaoLarkPosItemprodFindterminalAPIRequest{
    return &TaobaoLarkPosItemprodFindterminalAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetApiMethodName() string {
    return "taobao.lark.pos.itemprod.findterminal"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 终端id
func (r *TaobaoLarkPosItemprodFindterminalAPIRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetDeviceId() string {
    return r._deviceId
}
// DeviceType Setter
// 终端类型
func (r *TaobaoLarkPosItemprodFindterminalAPIRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetDeviceType() string {
    return r._deviceType
}
// CreateUser Setter
// 912874323429834
func (r *TaobaoLarkPosItemprodFindterminalAPIRequest) SetCreateUser(_createUser string) error {
    r._createUser = _createUser
    r.Set("create_user", _createUser)
    return nil
}

// CreateUser Getter
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetCreateUser() string {
    return r._createUser
}
// LeaseCode Setter
// 租户编码
func (r *TaobaoLarkPosItemprodFindterminalAPIRequest) SetLeaseCode(_leaseCode string) error {
    r._leaseCode = _leaseCode
    r.Set("lease_code", _leaseCode)
    return nil
}

// LeaseCode Getter
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetLeaseCode() string {
    return r._leaseCode
}
// CinemaId Setter
// 影城id
func (r *TaobaoLarkPosItemprodFindterminalAPIRequest) SetCinemaId(_cinemaId string) error {
    r._cinemaId = _cinemaId
    r.Set("cinema_id", _cinemaId)
    return nil
}

// CinemaId Getter
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetCinemaId() string {
    return r._cinemaId
}
// CinemaName Setter
// 影城名称
func (r *TaobaoLarkPosItemprodFindterminalAPIRequest) SetCinemaName(_cinemaName string) error {
    r._cinemaName = _cinemaName
    r.Set("cinema_name", _cinemaName)
    return nil
}

// CinemaName Getter
func (r TaobaoLarkPosItemprodFindterminalAPIRequest) GetCinemaName() string {
    return r._cinemaName
}