package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租车回传预约信息 API请求
tmall.car.lease.reserve

租赁公司回传预约到店信息
*/
type TmallCarLeaseReserveRequest struct {
    model.Params
    // 买家id
    buyerId   int64
    // 订单id
    orderId   int64
    // 文案
    text   string
    // 车架号
    vin   string
    // 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
    flag   int64
    // 买家昵称
    buyerNick   string
}

// 初始化TmallCarLeaseReserveRequest对象
func NewTmallCarLeaseReserveRequest() *TmallCarLeaseReserveRequest{
    return &TmallCarLeaseReserveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseReserveRequest) GetApiMethodName() string {
    return "tmall.car.lease.reserve"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseReserveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerId Setter
// 买家id
func (r *TmallCarLeaseReserveRequest) SetBuyerId(buyerId int64) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r TmallCarLeaseReserveRequest) GetBuyerId() int64 {
    return r.buyerId
}
// OrderId Setter
// 订单id
func (r *TmallCarLeaseReserveRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseReserveRequest) GetOrderId() int64 {
    return r.orderId
}
// Text Setter
// 文案
func (r *TmallCarLeaseReserveRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

// Text Getter
func (r TmallCarLeaseReserveRequest) GetText() string {
    return r.text
}
// Vin Setter
// 车架号
func (r *TmallCarLeaseReserveRequest) SetVin(vin string) error {
    r.vin = vin
    r.Set("vin", vin)
    return nil
}

// Vin Getter
func (r TmallCarLeaseReserveRequest) GetVin() string {
    return r.vin
}
// Flag Setter
// 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
func (r *TmallCarLeaseReserveRequest) SetFlag(flag int64) error {
    r.flag = flag
    r.Set("flag", flag)
    return nil
}

// Flag Getter
func (r TmallCarLeaseReserveRequest) GetFlag() int64 {
    return r.flag
}
// BuyerNick Setter
// 买家昵称
func (r *TmallCarLeaseReserveRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r TmallCarLeaseReserveRequest) GetBuyerNick() string {
    return r.buyerNick
}
