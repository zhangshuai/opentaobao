package idleisv

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼服务商订单价格修改接口 API请求
alibaba.idle.isv.order.adjustprice

闲鱼用户通过授权的服务商修改订单价格和邮费
*/
type AlibabaIdleIsvOrderAdjustpriceRequest struct {
    model.Params
    // 输入参数
    paramAdjustOrderPrice   *IsvAdjustOrderPriceDto
}

// 初始化AlibabaIdleIsvOrderAdjustpriceRequest对象
func NewAlibabaIdleIsvOrderAdjustpriceRequest() *AlibabaIdleIsvOrderAdjustpriceRequest{
    return &AlibabaIdleIsvOrderAdjustpriceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvOrderAdjustpriceRequest) GetApiMethodName() string {
    return "alibaba.idle.isv.order.adjustprice"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvOrderAdjustpriceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAdjustOrderPrice Setter
// 输入参数
func (r *AlibabaIdleIsvOrderAdjustpriceRequest) SetParamAdjustOrderPrice(paramAdjustOrderPrice *IsvAdjustOrderPriceDto) error {
    r.paramAdjustOrderPrice = paramAdjustOrderPrice
    r.Set("param_adjust_order_price", paramAdjustOrderPrice)
    return nil
}

// ParamAdjustOrderPrice Getter
func (r AlibabaIdleIsvOrderAdjustpriceRequest) GetParamAdjustOrderPrice() *IsvAdjustOrderPriceDto {
    return r.paramAdjustOrderPrice
}