package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseprojectactivitysyncAPIRequest 电商券数据同步 API请求
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
type AlibabaalihouseprojectactivitysyncAPIRequest struct {
	model.Params
	// 数据
	_businessActivityDataDto *BusinessActivityDataDto
}

// NewAlibabaalihouseprojectactivitysyncRequest 初始化AlibabaalihouseprojectactivitysyncAPIRequest对象
func NewAlibabaalihouseprojectactivitysyncRequest() *AlibabaalihouseprojectactivitysyncAPIRequest {
	return &AlibabaalihouseprojectactivitysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseprojectactivitysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.project.activity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseprojectactivitysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseprojectactivitysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusinessActivityDataDto is BusinessActivityDataDto Setter
// 数据
func (r *AlibabaalihouseprojectactivitysyncAPIRequest) SetBusinessActivityDataDto(_businessActivityDataDto *BusinessActivityDataDto) error {
	r._businessActivityDataDto = _businessActivityDataDto
	r.Set("business_activity_data_dto", _businessActivityDataDto)
	return nil
}

// GetBusinessActivityDataDto BusinessActivityDataDto Getter
func (r AlibabaalihouseprojectactivitysyncAPIRequest) GetBusinessActivityDataDto() *BusinessActivityDataDto {
	return r._businessActivityDataDto
}
