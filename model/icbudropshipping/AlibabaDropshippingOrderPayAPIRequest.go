package icbudropshipping

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
alibaba dropshipping 支付代扣 API请求
alibaba.dropshipping.order.pay

alibaba dropshipping 支付代扣
*/
type AlibabaDropshippingOrderPayAPIRequest struct {
    model.Params
    // request model
    _paramOrderPayRequest   *OrderPayRequest
}

// 初始化AlibabaDropshippingOrderPayAPIRequest对象
func NewAlibabaDropshippingOrderPayRequest() *AlibabaDropshippingOrderPayAPIRequest{
    return &AlibabaDropshippingOrderPayAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDropshippingOrderPayAPIRequest) GetApiMethodName() string {
    return "alibaba.dropshipping.order.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDropshippingOrderPayAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOrderPayRequest Setter
// request model
func (r *AlibabaDropshippingOrderPayAPIRequest) SetParamOrderPayRequest(_paramOrderPayRequest *OrderPayRequest) error {
    r._paramOrderPayRequest = _paramOrderPayRequest
    r.Set("param_order_pay_request", _paramOrderPayRequest)
    return nil
}

// ParamOrderPayRequest Getter
func (r AlibabaDropshippingOrderPayAPIRequest) GetParamOrderPayRequest() *OrderPayRequest {
    return r._paramOrderPayRequest
}