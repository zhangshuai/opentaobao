package aetask

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressInteractiveTaskDeliveryQueryAPIRequest AE互动任务投放 API请求
// aliexpress.interactive.task.delivery.query
//
// 将内部配置好的任务，如浏览商品，店铺投放给外部ISV
type AliexpressInteractiveTaskDeliveryQueryAPIRequest struct {
	model.Params
	// 返回结果
	_requestDto *QueryDeliveryRequestDto
}

// NewAliexpressInteractiveTaskDeliveryQueryRequest 初始化AliexpressInteractiveTaskDeliveryQueryAPIRequest对象
func NewAliexpressInteractiveTaskDeliveryQueryRequest() *AliexpressInteractiveTaskDeliveryQueryAPIRequest {
	return &AliexpressInteractiveTaskDeliveryQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressInteractiveTaskDeliveryQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.interactive.task.delivery.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressInteractiveTaskDeliveryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressInteractiveTaskDeliveryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestDto is RequestDto Setter
// 返回结果
func (r *AliexpressInteractiveTaskDeliveryQueryAPIRequest) SetRequestDto(_requestDto *QueryDeliveryRequestDto) error {
	r._requestDto = _requestDto
	r.Set("request_dto", _requestDto)
	return nil
}

// GetRequestDto RequestDto Getter
func (r AliexpressInteractiveTaskDeliveryQueryAPIRequest) GetRequestDto() *QueryDeliveryRequestDto {
	return r._requestDto
}
