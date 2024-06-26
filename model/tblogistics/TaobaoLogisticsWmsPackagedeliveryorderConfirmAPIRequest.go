package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest 包裹出库单确认 API请求
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
type TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest struct {
	model.Params
	// 请求
	_confirmPackageDeliveryOrderRequest *ConfirmPackageOrderRequest
}

// NewTaobaoLogisticsWmsPackagedeliveryorderConfirmRequest 初始化TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest对象
func NewTaobaoLogisticsWmsPackagedeliveryorderConfirmRequest() *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest {
	return &TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) Reset() {
	r._confirmPackageDeliveryOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.wms.packagedeliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmPackageDeliveryOrderRequest is ConfirmPackageDeliveryOrderRequest Setter
// 请求
func (r *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) SetConfirmPackageDeliveryOrderRequest(_confirmPackageDeliveryOrderRequest *ConfirmPackageOrderRequest) error {
	r._confirmPackageDeliveryOrderRequest = _confirmPackageDeliveryOrderRequest
	r.Set("confirm_package_delivery_order_request", _confirmPackageDeliveryOrderRequest)
	return nil
}

// GetConfirmPackageDeliveryOrderRequest ConfirmPackageDeliveryOrderRequest Getter
func (r TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) GetConfirmPackageDeliveryOrderRequest() *ConfirmPackageOrderRequest {
	return r._confirmPackageDeliveryOrderRequest
}

var poolTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsWmsPackagedeliveryorderConfirmRequest()
	},
}

// GetTaobaoLogisticsWmsPackagedeliveryorderConfirmRequest 从 sync.Pool 获取 TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest
func GetTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest() *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest {
	return poolTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest.Get().(*TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest)
}

// ReleaseTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest 将 TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest(v *TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest.Put(v)
}
