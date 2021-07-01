package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单 API请求
youku.ott.pay.order.queryorder

通过订单号查询订单信息
*/
type YoukuOttPayOrderQueryorderAPIRequest struct {
    model.Params
    // 订单号
    _orderNo   string
}

// 初始化YoukuOttPayOrderQueryorderAPIRequest对象
func NewYoukuOttPayOrderQueryorderRequest() *YoukuOttPayOrderQueryorderAPIRequest{
    return &YoukuOttPayOrderQueryorderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryorderAPIRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.queryorder"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryorderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 订单号
func (r *YoukuOttPayOrderQueryorderAPIRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r YoukuOttPayOrderQueryorderAPIRequest) GetOrderNo() string {
    return r._orderNo
}