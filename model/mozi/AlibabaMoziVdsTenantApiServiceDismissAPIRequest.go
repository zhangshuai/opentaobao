package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceDismissAPIRequest MOZI解除组织主管服务 API请求
// alibaba.mozi.vds.tenant.api.service.dismiss
//
// 解除组织主管
type AlibabaMoziVdsTenantApiServiceDismissAPIRequest struct {
	model.Params
	// 第一个入参
	_par0 *DismissOrganizationSupervisorRequest
}

// NewAlibabaMoziVdsTenantApiServiceDismissRequest 初始化AlibabaMoziVdsTenantApiServiceDismissAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceDismissRequest() *AlibabaMoziVdsTenantApiServiceDismissAPIRequest {
	return &AlibabaMoziVdsTenantApiServiceDismissAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziVdsTenantApiServiceDismissAPIRequest) Reset() {
	r._par0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceDismissAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.dismiss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceDismissAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServiceDismissAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 第一个入参
func (r *AlibabaMoziVdsTenantApiServiceDismissAPIRequest) SetPar0(_par0 *DismissOrganizationSupervisorRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceDismissAPIRequest) GetPar0() *DismissOrganizationSupervisorRequest {
	return r._par0
}

var poolAlibabaMoziVdsTenantApiServiceDismissAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziVdsTenantApiServiceDismissRequest()
	},
}

// GetAlibabaMoziVdsTenantApiServiceDismissRequest 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceDismissAPIRequest
func GetAlibabaMoziVdsTenantApiServiceDismissAPIRequest() *AlibabaMoziVdsTenantApiServiceDismissAPIRequest {
	return poolAlibabaMoziVdsTenantApiServiceDismissAPIRequest.Get().(*AlibabaMoziVdsTenantApiServiceDismissAPIRequest)
}

// ReleaseAlibabaMoziVdsTenantApiServiceDismissAPIRequest 将 AlibabaMoziVdsTenantApiServiceDismissAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceDismissAPIRequest(v *AlibabaMoziVdsTenantApiServiceDismissAPIRequest) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceDismissAPIRequest.Put(v)
}
