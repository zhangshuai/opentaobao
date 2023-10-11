package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCcoSelfCoordinateBreakOrderAPIRequest 天猫国际直购供应商毁单通知 API请求
// taobao.cco.self.coordinate.break.order
//
// 天猫国际直购供应商毁单通知
type TaobaoCcoSelfCoordinateBreakOrderAPIRequest struct {
	model.Params
	// 毁单请求参数
	_supplierBreakOrderRequest *SupplierBreakOrderRequest
}

// NewTaobaoCcoSelfCoordinateBreakOrderRequest 初始化TaobaoCcoSelfCoordinateBreakOrderAPIRequest对象
func NewTaobaoCcoSelfCoordinateBreakOrderRequest() *TaobaoCcoSelfCoordinateBreakOrderAPIRequest {
	return &TaobaoCcoSelfCoordinateBreakOrderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCcoSelfCoordinateBreakOrderAPIRequest) GetApiMethodName() string {
	return "taobao.cco.self.coordinate.break.order"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCcoSelfCoordinateBreakOrderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCcoSelfCoordinateBreakOrderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierBreakOrderRequest is SupplierBreakOrderRequest Setter
// 毁单请求参数
func (r *TaobaoCcoSelfCoordinateBreakOrderAPIRequest) SetSupplierBreakOrderRequest(_supplierBreakOrderRequest *SupplierBreakOrderRequest) error {
	r._supplierBreakOrderRequest = _supplierBreakOrderRequest
	r.Set("supplier_break_order_request", _supplierBreakOrderRequest)
	return nil
}

// GetSupplierBreakOrderRequest SupplierBreakOrderRequest Getter
func (r TaobaoCcoSelfCoordinateBreakOrderAPIRequest) GetSupplierBreakOrderRequest() *SupplierBreakOrderRequest {
	return r._supplierBreakOrderRequest
}
