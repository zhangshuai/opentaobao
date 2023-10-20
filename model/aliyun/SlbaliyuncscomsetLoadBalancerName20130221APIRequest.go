package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComSetLoadBalancerName20130221APIRequest 配置LoadBalancer的别名。 API请求
// slb.aliyuncs.com.SetLoadBalancerName.2013-02-21
//
// 配置LoadBalancer的别名。
type SlbAliyuncsComSetLoadBalancerName20130221APIRequest struct {
	model.Params
	// loadBalancerId
	_loadBalancerId string
	// loadBalancerName
	_loadBalancerName string
}

// NewSlbAliyuncsComSetLoadBalancerName20130221Request 初始化SlbAliyuncsComSetLoadBalancerName20130221APIRequest对象
func NewSlbAliyuncsComSetLoadBalancerName20130221Request() *SlbAliyuncsComSetLoadBalancerName20130221APIRequest {
	return &SlbAliyuncsComSetLoadBalancerName20130221APIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r SlbAliyuncsComSetLoadBalancerName20130221APIRequest) GetApiMethodName() string {
	return "slb.aliyuncs.com.SetLoadBalancerName.2013-02-21"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r SlbAliyuncsComSetLoadBalancerName20130221APIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r SlbAliyuncsComSetLoadBalancerName20130221APIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLoadBalancerId is LoadBalancerId Setter
// loadBalancerId
func (r *SlbAliyuncsComSetLoadBalancerName20130221APIRequest) SetLoadBalancerId(_loadBalancerId string) error {
	r._loadBalancerId = _loadBalancerId
	r.Set("loadBalancerId", _loadBalancerId)
	return nil
}

// GetLoadBalancerId LoadBalancerId Getter
func (r SlbAliyuncsComSetLoadBalancerName20130221APIRequest) GetLoadBalancerId() string {
	return r._loadBalancerId
}

// SetLoadBalancerName is LoadBalancerName Setter
// loadBalancerName
func (r *SlbAliyuncsComSetLoadBalancerName20130221APIRequest) SetLoadBalancerName(_loadBalancerName string) error {
	r._loadBalancerName = _loadBalancerName
	r.Set("loadBalancerName", _loadBalancerName)
	return nil
}

// GetLoadBalancerName LoadBalancerName Getter
func (r SlbAliyuncsComSetLoadBalancerName20130221APIRequest) GetLoadBalancerName() string {
	return r._loadBalancerName
}
