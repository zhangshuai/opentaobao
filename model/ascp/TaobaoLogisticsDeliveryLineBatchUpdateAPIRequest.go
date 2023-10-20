package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest 线路能力创建/更新 API请求
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
type TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest struct {
	model.Params
	// 线路能力创建/更新入参
	_deliveryLineBatchUpdateRequest *DeliveryLineBatchUpdateRequest
}

// NewTaobaoLogisticsDeliveryLineBatchUpdateRequest 初始化TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest对象
func NewTaobaoLogisticsDeliveryLineBatchUpdateRequest() *TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest {
	return &TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.delivery.line.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryLineBatchUpdateRequest is DeliveryLineBatchUpdateRequest Setter
// 线路能力创建/更新入参
func (r *TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest) SetDeliveryLineBatchUpdateRequest(_deliveryLineBatchUpdateRequest *DeliveryLineBatchUpdateRequest) error {
	r._deliveryLineBatchUpdateRequest = _deliveryLineBatchUpdateRequest
	r.Set("delivery_line_batch_update_request", _deliveryLineBatchUpdateRequest)
	return nil
}

// GetDeliveryLineBatchUpdateRequest DeliveryLineBatchUpdateRequest Getter
func (r TaobaoLogisticsDeliveryLineBatchUpdateAPIRequest) GetDeliveryLineBatchUpdateRequest() *DeliveryLineBatchUpdateRequest {
	return r._deliveryLineBatchUpdateRequest
}
