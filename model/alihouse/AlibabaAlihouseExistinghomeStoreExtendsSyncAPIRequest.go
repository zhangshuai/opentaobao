package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest 门店扩展信息变更 API请求
// alibaba.alihouse.existinghome.store.extends.sync
//
// 门店扩展信息变更
type AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest struct {
	model.Params
	// dto
	_storeExtendsInfoDto *StoreExtendsInfoDto
}

// NewAlibabaAlihouseExistinghomeStoreExtendsSyncRequest 初始化AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest对象
func NewAlibabaAlihouseExistinghomeStoreExtendsSyncRequest() *AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest {
	return &AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.extends.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreExtendsInfoDto is StoreExtendsInfoDto Setter
// dto
func (r *AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest) SetStoreExtendsInfoDto(_storeExtendsInfoDto *StoreExtendsInfoDto) error {
	r._storeExtendsInfoDto = _storeExtendsInfoDto
	r.Set("store_extends_info_dto", _storeExtendsInfoDto)
	return nil
}

// GetStoreExtendsInfoDto StoreExtendsInfoDto Getter
func (r AlibabaAlihouseExistinghomeStoreExtendsSyncAPIRequest) GetStoreExtendsInfoDto() *StoreExtendsInfoDto {
	return r._storeExtendsInfoDto
}
