package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest 获取设备状态信息 API请求
// taobao.ailab.aicloud.top.device.statusinfo.get
//
// 获取设备状态信息
type TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest struct {
	model.Params
	// 三方用户id或淘宝openId
	_originUserId string
	// 账号秘钥
	_schemaKey string
	// 三方传extUser，淘宝传openTaoBaoUser
	_userType string
	// 设备id
	_deviceId string
}

// NewTaobaoAilabAicloudTopDeviceStatusinfoGetRequest 初始化TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceStatusinfoGetRequest() *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest {
	return &TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.statusinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) SetOriginUserId(_originUserId string) error {
	r._originUserId = _originUserId
	r.Set("origin_user_id", _originUserId)
	return nil
}

// Get OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) GetOriginUserId() string {
	return r._originUserId
}

// Set is SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// Get SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// Set is UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// Get UserType Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) GetUserType() string {
	return r._userType
}

// Set is DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// Get DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceStatusinfoGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}
