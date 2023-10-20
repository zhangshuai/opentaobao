package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscusercenterinfoqueryAPIRequest 查询授权的用户信息 API请求
// alibaba.alsc.user.center.info.query
//
// 获取授权的饿了么用户信息
type AlibabaalscusercenterinfoqueryAPIRequest struct {
	model.Params
	// 请求模型
	_alscUserQueryOpenRequest *AlscUserQueryOpenRequest
}

// NewAlibabaalscusercenterinfoqueryRequest 初始化AlibabaalscusercenterinfoqueryAPIRequest对象
func NewAlibabaalscusercenterinfoqueryRequest() *AlibabaalscusercenterinfoqueryAPIRequest {
	return &AlibabaalscusercenterinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscusercenterinfoqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.user.center.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscusercenterinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscusercenterinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlscUserQueryOpenRequest is AlscUserQueryOpenRequest Setter
// 请求模型
func (r *AlibabaalscusercenterinfoqueryAPIRequest) SetAlscUserQueryOpenRequest(_alscUserQueryOpenRequest *AlscUserQueryOpenRequest) error {
	r._alscUserQueryOpenRequest = _alscUserQueryOpenRequest
	r.Set("alsc_user_query_open_request", _alscUserQueryOpenRequest)
	return nil
}

// GetAlscUserQueryOpenRequest AlscUserQueryOpenRequest Getter
func (r AlibabaalscusercenterinfoqueryAPIRequest) GetAlscUserQueryOpenRequest() *AlscUserQueryOpenRequest {
	return r._alscUserQueryOpenRequest
}