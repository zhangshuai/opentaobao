package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POS收银成功后订单同步 API请求
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步
*/
type AlibabaMjOcPayRequest struct {
    model.Params
    // 订单数据
    posOrder   *PosOrderDto
}

// 初始化AlibabaMjOcPayRequest对象
func NewAlibabaMjOcPayRequest() *AlibabaMjOcPayRequest{
    return &AlibabaMjOcPayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcPayRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.pay"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PosOrder Setter
// 订单数据
func (r *AlibabaMjOcPayRequest) SetPosOrder(posOrder *PosOrderDto) error {
    r.posOrder = posOrder
    r.Set("pos_order", posOrder)
    return nil
}

// PosOrder Getter
func (r AlibabaMjOcPayRequest) GetPosOrder() *PosOrderDto {
    return r.posOrder
}
