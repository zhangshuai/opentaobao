package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkRexoutDeviceInfoGetAPIRequest 获取设备详情-外部对接 API请求
// wdk.rexout.device.info.get
//
// 获取设备详情-外部对接
type WdkRexoutDeviceInfoGetAPIRequest struct {
	model.Params
	// ak
	_accessKey string
	// 签名
	_accessSign string
	// 设备ID
	_uuid string
}

// NewWdkRexoutDeviceInfoGetRequest 初始化WdkRexoutDeviceInfoGetAPIRequest对象
func NewWdkRexoutDeviceInfoGetRequest() *WdkRexoutDeviceInfoGetAPIRequest {
	return &WdkRexoutDeviceInfoGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkRexoutDeviceInfoGetAPIRequest) Reset() {
	r._accessKey = ""
	r._accessSign = ""
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkRexoutDeviceInfoGetAPIRequest) GetApiMethodName() string {
	return "wdk.rexout.device.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkRexoutDeviceInfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkRexoutDeviceInfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// ak
func (r *WdkRexoutDeviceInfoGetAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r WdkRexoutDeviceInfoGetAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetAccessSign is AccessSign Setter
// 签名
func (r *WdkRexoutDeviceInfoGetAPIRequest) SetAccessSign(_accessSign string) error {
	r._accessSign = _accessSign
	r.Set("access_sign", _accessSign)
	return nil
}

// GetAccessSign AccessSign Getter
func (r WdkRexoutDeviceInfoGetAPIRequest) GetAccessSign() string {
	return r._accessSign
}

// SetUuid is Uuid Setter
// 设备ID
func (r *WdkRexoutDeviceInfoGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r WdkRexoutDeviceInfoGetAPIRequest) GetUuid() string {
	return r._uuid
}

var poolWdkRexoutDeviceInfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkRexoutDeviceInfoGetRequest()
	},
}

// GetWdkRexoutDeviceInfoGetRequest 从 sync.Pool 获取 WdkRexoutDeviceInfoGetAPIRequest
func GetWdkRexoutDeviceInfoGetAPIRequest() *WdkRexoutDeviceInfoGetAPIRequest {
	return poolWdkRexoutDeviceInfoGetAPIRequest.Get().(*WdkRexoutDeviceInfoGetAPIRequest)
}

// ReleaseWdkRexoutDeviceInfoGetAPIRequest 将 WdkRexoutDeviceInfoGetAPIRequest 放入 sync.Pool
func ReleaseWdkRexoutDeviceInfoGetAPIRequest(v *WdkRexoutDeviceInfoGetAPIRequest) {
	v.Reset()
	poolWdkRexoutDeviceInfoGetAPIRequest.Put(v)
}
