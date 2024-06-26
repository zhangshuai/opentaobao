package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreSyncAPIRequest 二手房标准门店数据同步 API请求
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
type AlibabaAlihouseExistinghomeStoreSyncAPIRequest struct {
	model.Params
	// 入参
	_companyStoreDto *CompanyStoreDto
}

// NewAlibabaAlihouseExistinghomeStoreSyncRequest 初始化AlibabaAlihouseExistinghomeStoreSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeStoreSyncRequest() *AlibabaAlihouseExistinghomeStoreSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeStoreSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeStoreSyncAPIRequest) Reset() {
	r._companyStoreDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeStoreSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeStoreSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeStoreSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCompanyStoreDto is CompanyStoreDto Setter
// 入参
func (r *AlibabaAlihouseExistinghomeStoreSyncAPIRequest) SetCompanyStoreDto(_companyStoreDto *CompanyStoreDto) error {
	r._companyStoreDto = _companyStoreDto
	r.Set("company_store_dto", _companyStoreDto)
	return nil
}

// GetCompanyStoreDto CompanyStoreDto Getter
func (r AlibabaAlihouseExistinghomeStoreSyncAPIRequest) GetCompanyStoreDto() *CompanyStoreDto {
	return r._companyStoreDto
}

var poolAlibabaAlihouseExistinghomeStoreSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeStoreSyncRequest()
	},
}

// GetAlibabaAlihouseExistinghomeStoreSyncRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreSyncAPIRequest
func GetAlibabaAlihouseExistinghomeStoreSyncAPIRequest() *AlibabaAlihouseExistinghomeStoreSyncAPIRequest {
	return poolAlibabaAlihouseExistinghomeStoreSyncAPIRequest.Get().(*AlibabaAlihouseExistinghomeStoreSyncAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeStoreSyncAPIRequest 将 AlibabaAlihouseExistinghomeStoreSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreSyncAPIRequest(v *AlibabaAlihouseExistinghomeStoreSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreSyncAPIRequest.Put(v)
}
