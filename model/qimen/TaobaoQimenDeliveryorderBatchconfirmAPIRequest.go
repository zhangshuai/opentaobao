package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchconfirmAPIRequest 发货单确认接口 API请求
// taobao.qimen.deliveryorder.batchconfirm
//
// taobao.qimen.deliveryorder.batchconfirm
type TaobaoQimenDeliveryorderBatchconfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchConfirmRequest
}

// NewTaobaoQimenDeliveryorderBatchconfirmRequest 初始化TaobaoQimenDeliveryorderBatchconfirmAPIRequest对象
func NewTaobaoQimenDeliveryorderBatchconfirmRequest() *TaobaoQimenDeliveryorderBatchconfirmAPIRequest {
	return &TaobaoQimenDeliveryorderBatchconfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenDeliveryorderBatchconfirmAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderBatchconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenDeliveryorderBatchconfirmAPIRequest) SetRequest(_request *DeliveryOrderBatchConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderBatchconfirmAPIRequest) GetRequest() *DeliveryOrderBatchConfirmRequest {
	return r._request
}

var poolTaobaoQimenDeliveryorderBatchconfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenDeliveryorderBatchconfirmRequest()
	},
}

// GetTaobaoQimenDeliveryorderBatchconfirmRequest 从 sync.Pool 获取 TaobaoQimenDeliveryorderBatchconfirmAPIRequest
func GetTaobaoQimenDeliveryorderBatchconfirmAPIRequest() *TaobaoQimenDeliveryorderBatchconfirmAPIRequest {
	return poolTaobaoQimenDeliveryorderBatchconfirmAPIRequest.Get().(*TaobaoQimenDeliveryorderBatchconfirmAPIRequest)
}

// ReleaseTaobaoQimenDeliveryorderBatchconfirmAPIRequest 将 TaobaoQimenDeliveryorderBatchconfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenDeliveryorderBatchconfirmAPIRequest(v *TaobaoQimenDeliveryorderBatchconfirmAPIRequest) {
	v.Reset()
	poolTaobaoQimenDeliveryorderBatchconfirmAPIRequest.Put(v)
}
