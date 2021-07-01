package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退款申请接口 API请求
taobao.bus.refundticketprice.set

汽车票代理商利用该接口申请退票
*/
type TaobaoBusRefundticketpriceSetAPIRequest struct {
    model.Params
    // 退票申请入参
    _offlineRefundTicketRq   *OfflineRefundTicketPriceRq
}

// 初始化TaobaoBusRefundticketpriceSetAPIRequest对象
func NewTaobaoBusRefundticketpriceSetRequest() *TaobaoBusRefundticketpriceSetAPIRequest{
    return &TaobaoBusRefundticketpriceSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetApiMethodName() string {
    return "taobao.bus.refundticketprice.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineRefundTicketRq Setter
// 退票申请入参
func (r *TaobaoBusRefundticketpriceSetAPIRequest) SetOfflineRefundTicketRq(_offlineRefundTicketRq *OfflineRefundTicketPriceRq) error {
    r._offlineRefundTicketRq = _offlineRefundTicketRq
    r.Set("offline_refund_ticket_rq", _offlineRefundTicketRq)
    return nil
}

// OfflineRefundTicketRq Getter
func (r TaobaoBusRefundticketpriceSetAPIRequest) GetOfflineRefundTicketRq() *OfflineRefundTicketPriceRq {
    return r._offlineRefundTicketRq
}