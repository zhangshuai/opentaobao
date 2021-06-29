package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单id获取单条订单详情 API请求
alibaba.alihealth.nr.trade.order.getorderdetail

阿里健康O2O，获取订单详情，修复组合商品价格精度问题
*/
type AlibabaAlihealthNrTradeOrderGetorderdetailRequest struct {
    model.Params
    // 淘宝订单ID
    orderId   int64
}

// 初始化AlibabaAlihealthNrTradeOrderGetorderdetailRequest对象
func NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest() *AlibabaAlihealthNrTradeOrderGetorderdetailRequest{
    return &AlibabaAlihealthNrTradeOrderGetorderdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrTradeOrderGetorderdetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.trade.order.getorderdetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrTradeOrderGetorderdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 淘宝订单ID
func (r *AlibabaAlihealthNrTradeOrderGetorderdetailRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthNrTradeOrderGetorderdetailRequest) GetOrderId() int64 {
    return r.orderId
}
