package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasscsupplyplatformserviceabilitydeleteAPIRequest 删除服务能力 API请求
// alibaba.ssc.supplyplatform.serviceability.delete
//
// 删除服务能力
type AlibabasscsupplyplatformserviceabilitydeleteAPIRequest struct {
	model.Params
	// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
	_providerType string
	// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
	_providerId int64
}

// NewAlibabasscsupplyplatformserviceabilitydeleteRequest 初始化AlibabasscsupplyplatformserviceabilitydeleteAPIRequest对象
func NewAlibabasscsupplyplatformserviceabilitydeleteRequest() *AlibabasscsupplyplatformserviceabilitydeleteAPIRequest {
	return &AlibabasscsupplyplatformserviceabilitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.supplyplatform.serviceability.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProviderType is ProviderType Setter
// 服务提供者类型。service_store 网点；worker 工人；supplier 服务商
func (r *AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) SetProviderType(_providerType string) error {
	r._providerType = _providerType
	r.Set("provider_type", _providerType)
	return nil
}

// GetProviderType ProviderType Getter
func (r AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) GetProviderType() string {
	return r._providerType
}

// SetProviderId is ProviderId Setter
// 服务提供者id。根据服务提供者类型填写相应的id，例如类型是网点，则填我们系统的网点id
func (r *AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) SetProviderId(_providerId int64) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlibabasscsupplyplatformserviceabilitydeleteAPIRequest) GetProviderId() int64 {
	return r._providerId
}
