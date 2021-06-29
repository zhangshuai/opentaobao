package alimsg

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
虚拟发货消息发送接口 API请求
alibaba.idle.order.msg.send

用户下单后服务商期望自动发货，该接口用于给用户发送文本消息，主要用于卡券类等虚拟商品场景
*/
type AlibabaIdleOrderMsgSendRequest struct {
    model.Params
    // 订单id
    orderId   string
    // 消息发送内容
    text   string
}

// 初始化AlibabaIdleOrderMsgSendRequest对象
func NewAlibabaIdleOrderMsgSendRequest() *AlibabaIdleOrderMsgSendRequest{
    return &AlibabaIdleOrderMsgSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleOrderMsgSendRequest) GetApiMethodName() string {
    return "alibaba.idle.order.msg.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleOrderMsgSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleOrderMsgSendRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleOrderMsgSendRequest) GetOrderId() string {
    return r.orderId
}
// Text Setter
// 消息发送内容
func (r *AlibabaIdleOrderMsgSendRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

// Text Getter
func (r AlibabaIdleOrderMsgSendRequest) GetText() string {
    return r.text
}