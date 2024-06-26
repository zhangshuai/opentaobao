package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStoreAuditQueryAPIRequest ISV查询门店审核状态 API请求
// alibaba.alihealth.dental.store.audit.query
//
// ISV查询门店审核状态
type AlibabaAlihealthDentalStoreAuditQueryAPIRequest struct {
	model.Params
	// 审核ID列表
	_storeAuditIds []string
}

// NewAlibabaAlihealthDentalStoreAuditQueryRequest 初始化AlibabaAlihealthDentalStoreAuditQueryAPIRequest对象
func NewAlibabaAlihealthDentalStoreAuditQueryRequest() *AlibabaAlihealthDentalStoreAuditQueryAPIRequest {
	return &AlibabaAlihealthDentalStoreAuditQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDentalStoreAuditQueryAPIRequest) Reset() {
	r._storeAuditIds = r._storeAuditIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.store.audit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreAuditIds is StoreAuditIds Setter
// 审核ID列表
func (r *AlibabaAlihealthDentalStoreAuditQueryAPIRequest) SetStoreAuditIds(_storeAuditIds []string) error {
	r._storeAuditIds = _storeAuditIds
	r.Set("store_audit_ids", _storeAuditIds)
	return nil
}

// GetStoreAuditIds StoreAuditIds Getter
func (r AlibabaAlihealthDentalStoreAuditQueryAPIRequest) GetStoreAuditIds() []string {
	return r._storeAuditIds
}

var poolAlibabaAlihealthDentalStoreAuditQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDentalStoreAuditQueryRequest()
	},
}

// GetAlibabaAlihealthDentalStoreAuditQueryRequest 从 sync.Pool 获取 AlibabaAlihealthDentalStoreAuditQueryAPIRequest
func GetAlibabaAlihealthDentalStoreAuditQueryAPIRequest() *AlibabaAlihealthDentalStoreAuditQueryAPIRequest {
	return poolAlibabaAlihealthDentalStoreAuditQueryAPIRequest.Get().(*AlibabaAlihealthDentalStoreAuditQueryAPIRequest)
}

// ReleaseAlibabaAlihealthDentalStoreAuditQueryAPIRequest 将 AlibabaAlihealthDentalStoreAuditQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDentalStoreAuditQueryAPIRequest(v *AlibabaAlihealthDentalStoreAuditQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDentalStoreAuditQueryAPIRequest.Put(v)
}
