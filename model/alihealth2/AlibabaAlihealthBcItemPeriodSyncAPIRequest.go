package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBcItemPeriodSyncAPIRequest 代销品效期同步 API请求
// alibaba.alihealth.bc.item.period.sync
//
// 代销品效期同步
type AlibabaAlihealthBcItemPeriodSyncAPIRequest struct {
	model.Params
	// 请求体
	_validityPeriodSyncReqDto *ValidityPeriodSyncReqDto
}

// NewAlibabaAlihealthBcItemPeriodSyncRequest 初始化AlibabaAlihealthBcItemPeriodSyncAPIRequest对象
func NewAlibabaAlihealthBcItemPeriodSyncRequest() *AlibabaAlihealthBcItemPeriodSyncAPIRequest {
	return &AlibabaAlihealthBcItemPeriodSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBcItemPeriodSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.bc.item.period.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBcItemPeriodSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBcItemPeriodSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetValidityPeriodSyncReqDto is ValidityPeriodSyncReqDto Setter
// 请求体
func (r *AlibabaAlihealthBcItemPeriodSyncAPIRequest) SetValidityPeriodSyncReqDto(_validityPeriodSyncReqDto *ValidityPeriodSyncReqDto) error {
	r._validityPeriodSyncReqDto = _validityPeriodSyncReqDto
	r.Set("validity_period_sync_req_dto", _validityPeriodSyncReqDto)
	return nil
}

// GetValidityPeriodSyncReqDto ValidityPeriodSyncReqDto Getter
func (r AlibabaAlihealthBcItemPeriodSyncAPIRequest) GetValidityPeriodSyncReqDto() *ValidityPeriodSyncReqDto {
	return r._validityPeriodSyncReqDto
}
