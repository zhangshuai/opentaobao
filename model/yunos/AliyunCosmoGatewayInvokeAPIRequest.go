package yunos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunCosmoGatewayInvokeAPIRequest alios cosmo服务调用 API请求
// aliyun.cosmo.gateway.invoke
//
// AliOS cosmo服务分发平台对外调用接口
type AliyunCosmoGatewayInvokeAPIRequest struct {
	model.Params
	// 请求上下文参数
	_context *RdamContext
	// 请求对象
	_rdamRequest *RdamGenericRequest
}

// NewAliyunCosmoGatewayInvokeRequest 初始化AliyunCosmoGatewayInvokeAPIRequest对象
func NewAliyunCosmoGatewayInvokeRequest() *AliyunCosmoGatewayInvokeAPIRequest {
	return &AliyunCosmoGatewayInvokeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunCosmoGatewayInvokeAPIRequest) Reset() {
	r._context = nil
	r._rdamRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunCosmoGatewayInvokeAPIRequest) GetApiMethodName() string {
	return "aliyun.cosmo.gateway.invoke"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunCosmoGatewayInvokeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunCosmoGatewayInvokeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContext is Context Setter
// 请求上下文参数
func (r *AliyunCosmoGatewayInvokeAPIRequest) SetContext(_context *RdamContext) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r AliyunCosmoGatewayInvokeAPIRequest) GetContext() *RdamContext {
	return r._context
}

// SetRdamRequest is RdamRequest Setter
// 请求对象
func (r *AliyunCosmoGatewayInvokeAPIRequest) SetRdamRequest(_rdamRequest *RdamGenericRequest) error {
	r._rdamRequest = _rdamRequest
	r.Set("rdam_request", _rdamRequest)
	return nil
}

// GetRdamRequest RdamRequest Getter
func (r AliyunCosmoGatewayInvokeAPIRequest) GetRdamRequest() *RdamGenericRequest {
	return r._rdamRequest
}

var poolAliyunCosmoGatewayInvokeAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunCosmoGatewayInvokeRequest()
	},
}

// GetAliyunCosmoGatewayInvokeRequest 从 sync.Pool 获取 AliyunCosmoGatewayInvokeAPIRequest
func GetAliyunCosmoGatewayInvokeAPIRequest() *AliyunCosmoGatewayInvokeAPIRequest {
	return poolAliyunCosmoGatewayInvokeAPIRequest.Get().(*AliyunCosmoGatewayInvokeAPIRequest)
}

// ReleaseAliyunCosmoGatewayInvokeAPIRequest 将 AliyunCosmoGatewayInvokeAPIRequest 放入 sync.Pool
func ReleaseAliyunCosmoGatewayInvokeAPIRequest(v *AliyunCosmoGatewayInvokeAPIRequest) {
	v.Reset()
	poolAliyunCosmoGatewayInvokeAPIRequest.Put(v)
}
