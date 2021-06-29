package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单关闭接口（快速退款） API请求
alitrip.travel.trade.close

卖家关单（快速退款接口），不支持二次预约商品的订单
*/
type AlitripTravelTradeCloseRequest struct {
    model.Params
    // 交易关闭原因
    closeReason   string
    // 子订单编号
    subOrderId   int64
    // 订单关闭原因描述
    reasonDesc   string
}

// 初始化AlitripTravelTradeCloseRequest对象
func NewAlitripTravelTradeCloseRequest() *AlitripTravelTradeCloseRequest{
    return &AlitripTravelTradeCloseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeCloseRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.close"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CloseReason Setter
// 交易关闭原因
func (r *AlitripTravelTradeCloseRequest) SetCloseReason(closeReason string) error {
    r.closeReason = closeReason
    r.Set("close_reason", closeReason)
    return nil
}

// CloseReason Getter
func (r AlitripTravelTradeCloseRequest) GetCloseReason() string {
    return r.closeReason
}
// SubOrderId Setter
// 子订单编号
func (r *AlitripTravelTradeCloseRequest) SetSubOrderId(subOrderId int64) error {
    r.subOrderId = subOrderId
    r.Set("sub_order_id", subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlitripTravelTradeCloseRequest) GetSubOrderId() int64 {
    return r.subOrderId
}
// ReasonDesc Setter
// 订单关闭原因描述
func (r *AlitripTravelTradeCloseRequest) SetReasonDesc(reasonDesc string) error {
    r.reasonDesc = reasonDesc
    r.Set("reason_desc", reasonDesc)
    return nil
}

// ReasonDesc Getter
func (r AlitripTravelTradeCloseRequest) GetReasonDesc() string {
    return r.reasonDesc
}