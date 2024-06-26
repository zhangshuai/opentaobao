package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseStorePunishInfoQueryAPIRequest 门店处罚信息查询 API请求
// alibaba.alihouse.store.punish.info.query
//
// 门店处罚信息查询
type AlibabaAlihouseStorePunishInfoQueryAPIRequest struct {
	model.Params
	// dto
	_queryStorePunishDto *QueryStorePunishDto
}

// NewAlibabaAlihouseStorePunishInfoQueryRequest 初始化AlibabaAlihouseStorePunishInfoQueryAPIRequest对象
func NewAlibabaAlihouseStorePunishInfoQueryRequest() *AlibabaAlihouseStorePunishInfoQueryAPIRequest {
	return &AlibabaAlihouseStorePunishInfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseStorePunishInfoQueryAPIRequest) Reset() {
	r._queryStorePunishDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseStorePunishInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.store.punish.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseStorePunishInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseStorePunishInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryStorePunishDto is QueryStorePunishDto Setter
// dto
func (r *AlibabaAlihouseStorePunishInfoQueryAPIRequest) SetQueryStorePunishDto(_queryStorePunishDto *QueryStorePunishDto) error {
	r._queryStorePunishDto = _queryStorePunishDto
	r.Set("query_store_punish_dto", _queryStorePunishDto)
	return nil
}

// GetQueryStorePunishDto QueryStorePunishDto Getter
func (r AlibabaAlihouseStorePunishInfoQueryAPIRequest) GetQueryStorePunishDto() *QueryStorePunishDto {
	return r._queryStorePunishDto
}

var poolAlibabaAlihouseStorePunishInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseStorePunishInfoQueryRequest()
	},
}

// GetAlibabaAlihouseStorePunishInfoQueryRequest 从 sync.Pool 获取 AlibabaAlihouseStorePunishInfoQueryAPIRequest
func GetAlibabaAlihouseStorePunishInfoQueryAPIRequest() *AlibabaAlihouseStorePunishInfoQueryAPIRequest {
	return poolAlibabaAlihouseStorePunishInfoQueryAPIRequest.Get().(*AlibabaAlihouseStorePunishInfoQueryAPIRequest)
}

// ReleaseAlibabaAlihouseStorePunishInfoQueryAPIRequest 将 AlibabaAlihouseStorePunishInfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseStorePunishInfoQueryAPIRequest(v *AlibabaAlihouseStorePunishInfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseStorePunishInfoQueryAPIRequest.Put(v)
}
