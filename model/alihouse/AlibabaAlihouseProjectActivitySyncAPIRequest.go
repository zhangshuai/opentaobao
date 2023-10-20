package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseProjectActivitySyncAPIRequest 电商券数据同步 API请求
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
type AlibabaAlihouseProjectActivitySyncAPIRequest struct {
	model.Params
	// 数据
	_businessActivityDataDto *BusinessActivityDataDto
}

// NewAlibabaAlihouseProjectActivitySyncRequest 初始化AlibabaAlihouseProjectActivitySyncAPIRequest对象
func NewAlibabaAlihouseProjectActivitySyncRequest() *AlibabaAlihouseProjectActivitySyncAPIRequest {
	return &AlibabaAlihouseProjectActivitySyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseProjectActivitySyncAPIRequest) Reset() {
	r._businessActivityDataDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.project.activity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusinessActivityDataDto is BusinessActivityDataDto Setter
// 数据
func (r *AlibabaAlihouseProjectActivitySyncAPIRequest) SetBusinessActivityDataDto(_businessActivityDataDto *BusinessActivityDataDto) error {
	r._businessActivityDataDto = _businessActivityDataDto
	r.Set("business_activity_data_dto", _businessActivityDataDto)
	return nil
}

// GetBusinessActivityDataDto BusinessActivityDataDto Getter
func (r AlibabaAlihouseProjectActivitySyncAPIRequest) GetBusinessActivityDataDto() *BusinessActivityDataDto {
	return r._businessActivityDataDto
}

var poolAlibabaAlihouseProjectActivitySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseProjectActivitySyncRequest()
	},
}

// GetAlibabaAlihouseProjectActivitySyncRequest 从 sync.Pool 获取 AlibabaAlihouseProjectActivitySyncAPIRequest
func GetAlibabaAlihouseProjectActivitySyncAPIRequest() *AlibabaAlihouseProjectActivitySyncAPIRequest {
	return poolAlibabaAlihouseProjectActivitySyncAPIRequest.Get().(*AlibabaAlihouseProjectActivitySyncAPIRequest)
}

// ReleaseAlibabaAlihouseProjectActivitySyncAPIRequest 将 AlibabaAlihouseProjectActivitySyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseProjectActivitySyncAPIRequest(v *AlibabaAlihouseProjectActivitySyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseProjectActivitySyncAPIRequest.Put(v)
}
