package cmns

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaDeviceGetAPIRequest 设备详情查询 API请求
// yunos.service.cmns.coa.device.get
//
// 第三方应用开发者调用此接口查询设备详情
type YunosServiceCmnsCoaDeviceGetAPIRequest struct {
	model.Params
	// 设备id类型,可以是uuid,imei,deviceToken,kp
	_type string
	// 设备id
	_value string
}

// NewYunosServiceCmnsCoaDeviceGetRequest 初始化YunosServiceCmnsCoaDeviceGetAPIRequest对象
func NewYunosServiceCmnsCoaDeviceGetRequest() *YunosServiceCmnsCoaDeviceGetAPIRequest {
	return &YunosServiceCmnsCoaDeviceGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosServiceCmnsCoaDeviceGetAPIRequest) Reset() {
	r._type = ""
	r._value = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosServiceCmnsCoaDeviceGetAPIRequest) GetApiMethodName() string {
	return "yunos.service.cmns.coa.device.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosServiceCmnsCoaDeviceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosServiceCmnsCoaDeviceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 设备id类型,可以是uuid,imei,deviceToken,kp
func (r *YunosServiceCmnsCoaDeviceGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r YunosServiceCmnsCoaDeviceGetAPIRequest) GetType() string {
	return r._type
}

// SetValue is Value Setter
// 设备id
func (r *YunosServiceCmnsCoaDeviceGetAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r YunosServiceCmnsCoaDeviceGetAPIRequest) GetValue() string {
	return r._value
}

var poolYunosServiceCmnsCoaDeviceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosServiceCmnsCoaDeviceGetRequest()
	},
}

// GetYunosServiceCmnsCoaDeviceGetRequest 从 sync.Pool 获取 YunosServiceCmnsCoaDeviceGetAPIRequest
func GetYunosServiceCmnsCoaDeviceGetAPIRequest() *YunosServiceCmnsCoaDeviceGetAPIRequest {
	return poolYunosServiceCmnsCoaDeviceGetAPIRequest.Get().(*YunosServiceCmnsCoaDeviceGetAPIRequest)
}

// ReleaseYunosServiceCmnsCoaDeviceGetAPIRequest 将 YunosServiceCmnsCoaDeviceGetAPIRequest 放入 sync.Pool
func ReleaseYunosServiceCmnsCoaDeviceGetAPIRequest(v *YunosServiceCmnsCoaDeviceGetAPIRequest) {
	v.Reset()
	poolYunosServiceCmnsCoaDeviceGetAPIRequest.Put(v)
}
