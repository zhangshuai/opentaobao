package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosPubadminCommonOperationAPIRequest 内部迎客松通用服务 API请求
// yunos.pubadmin.common.operation
//
// 内部迎客松通用服务
type YunosPubadminCommonOperationAPIRequest struct {
	model.Params
	// 接口名
	_interfaceName string
	// 方法名
	_methodName string
	// 入参json串
	_parameter string
}

// NewYunosPubadminCommonOperationRequest 初始化YunosPubadminCommonOperationAPIRequest对象
func NewYunosPubadminCommonOperationRequest() *YunosPubadminCommonOperationAPIRequest {
	return &YunosPubadminCommonOperationAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosPubadminCommonOperationAPIRequest) Reset() {
	r._interfaceName = ""
	r._methodName = ""
	r._parameter = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosPubadminCommonOperationAPIRequest) GetApiMethodName() string {
	return "yunos.pubadmin.common.operation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosPubadminCommonOperationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosPubadminCommonOperationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInterfaceName is InterfaceName Setter
// 接口名
func (r *YunosPubadminCommonOperationAPIRequest) SetInterfaceName(_interfaceName string) error {
	r._interfaceName = _interfaceName
	r.Set("interface_name", _interfaceName)
	return nil
}

// GetInterfaceName InterfaceName Getter
func (r YunosPubadminCommonOperationAPIRequest) GetInterfaceName() string {
	return r._interfaceName
}

// SetMethodName is MethodName Setter
// 方法名
func (r *YunosPubadminCommonOperationAPIRequest) SetMethodName(_methodName string) error {
	r._methodName = _methodName
	r.Set("method_name", _methodName)
	return nil
}

// GetMethodName MethodName Getter
func (r YunosPubadminCommonOperationAPIRequest) GetMethodName() string {
	return r._methodName
}

// SetParameter is Parameter Setter
// 入参json串
func (r *YunosPubadminCommonOperationAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r YunosPubadminCommonOperationAPIRequest) GetParameter() string {
	return r._parameter
}

var poolYunosPubadminCommonOperationAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosPubadminCommonOperationRequest()
	},
}

// GetYunosPubadminCommonOperationRequest 从 sync.Pool 获取 YunosPubadminCommonOperationAPIRequest
func GetYunosPubadminCommonOperationAPIRequest() *YunosPubadminCommonOperationAPIRequest {
	return poolYunosPubadminCommonOperationAPIRequest.Get().(*YunosPubadminCommonOperationAPIRequest)
}

// ReleaseYunosPubadminCommonOperationAPIRequest 将 YunosPubadminCommonOperationAPIRequest 放入 sync.Pool
func ReleaseYunosPubadminCommonOperationAPIRequest(v *YunosPubadminCommonOperationAPIRequest) {
	v.Reset()
	poolYunosPubadminCommonOperationAPIRequest.Put(v)
}
