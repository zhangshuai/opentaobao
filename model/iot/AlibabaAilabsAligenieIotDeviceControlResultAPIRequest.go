package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieIotDeviceControlResultAPIRequest 设备控制结果 API请求
// alibaba.ailabs.aligenie.iot.device.control.result
//
// 智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
type AlibabaAilabsAligenieIotDeviceControlResultAPIRequest struct {
	model.Params
	// 请求token
	_requestToken string
	// 设备ID
	_deviceId string
	// 厂商执行返回ackCode
	_ackCode string
	// 操作类型 1：控制操作 0:查询
	_type int64
	// 控制成功
	_control bool
}

// NewAlibabaAilabsAligenieIotDeviceControlResultRequest 初始化AlibabaAilabsAligenieIotDeviceControlResultAPIRequest对象
func NewAlibabaAilabsAligenieIotDeviceControlResultRequest() *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest {
	return &AlibabaAilabsAligenieIotDeviceControlResultAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) Reset() {
	r._requestToken = ""
	r._deviceId = ""
	r._ackCode = ""
	r._type = 0
	r._control = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.iot.device.control.result"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestToken is RequestToken Setter
// 请求token
func (r *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) SetRequestToken(_requestToken string) error {
	r._requestToken = _requestToken
	r.Set("request_token", _requestToken)
	return nil
}

// GetRequestToken RequestToken Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetRequestToken() string {
	return r._requestToken
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetAckCode is AckCode Setter
// 厂商执行返回ackCode
func (r *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) SetAckCode(_ackCode string) error {
	r._ackCode = _ackCode
	r.Set("ack_code", _ackCode)
	return nil
}

// GetAckCode AckCode Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetAckCode() string {
	return r._ackCode
}

// SetType is Type Setter
// 操作类型 1：控制操作 0:查询
func (r *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetType() int64 {
	return r._type
}

// SetControl is Control Setter
// 控制成功
func (r *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) SetControl(_control bool) error {
	r._control = _control
	r.Set("control", _control)
	return nil
}

// GetControl Control Getter
func (r AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) GetControl() bool {
	return r._control
}

var poolAlibabaAilabsAligenieIotDeviceControlResultAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsAligenieIotDeviceControlResultRequest()
	},
}

// GetAlibabaAilabsAligenieIotDeviceControlResultRequest 从 sync.Pool 获取 AlibabaAilabsAligenieIotDeviceControlResultAPIRequest
func GetAlibabaAilabsAligenieIotDeviceControlResultAPIRequest() *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest {
	return poolAlibabaAilabsAligenieIotDeviceControlResultAPIRequest.Get().(*AlibabaAilabsAligenieIotDeviceControlResultAPIRequest)
}

// ReleaseAlibabaAilabsAligenieIotDeviceControlResultAPIRequest 将 AlibabaAilabsAligenieIotDeviceControlResultAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsAligenieIotDeviceControlResultAPIRequest(v *AlibabaAilabsAligenieIotDeviceControlResultAPIRequest) {
	v.Reset()
	poolAlibabaAilabsAligenieIotDeviceControlResultAPIRequest.Put(v)
}
