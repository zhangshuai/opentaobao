package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道卖品落单 API请求
taobao.lark.iot.order.confirmorder

云智对接无人超市，接收无人超市订单信息
*/
type TaobaoLarkIotOrderConfirmorderRequest struct {
    model.Params
    // 渠道编码
    channelCode   string
    // 影院内码
    cinemaLinkId   string
    // 外部订单号
    outGoodsOrderId   string
    // 工作站id
    workstationId   string
    // 工作站名称
    workstationName   string
    // 支付方式
    paymentList   string
    // 优惠列表
    promotionList   string
    // 卖品列表
    goodsList   string
    // 手机号
    mobile   string
    // 管理员
    operatorUserId   string
}

// 初始化TaobaoLarkIotOrderConfirmorderRequest对象
func NewTaobaoLarkIotOrderConfirmorderRequest() *TaobaoLarkIotOrderConfirmorderRequest{
    return &TaobaoLarkIotOrderConfirmorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderConfirmorderRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.confirmorder"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderConfirmorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetChannelCode(channelCode string) error {
    r.channelCode = channelCode
    r.Set("channel_code", channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetChannelCode() string {
    return r.channelCode
}
// CinemaLinkId Setter
// 影院内码
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetCinemaLinkId(cinemaLinkId string) error {
    r.cinemaLinkId = cinemaLinkId
    r.Set("cinema_link_id", cinemaLinkId)
    return nil
}

// CinemaLinkId Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetCinemaLinkId() string {
    return r.cinemaLinkId
}
// OutGoodsOrderId Setter
// 外部订单号
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetOutGoodsOrderId(outGoodsOrderId string) error {
    r.outGoodsOrderId = outGoodsOrderId
    r.Set("out_goods_order_id", outGoodsOrderId)
    return nil
}

// OutGoodsOrderId Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetOutGoodsOrderId() string {
    return r.outGoodsOrderId
}
// WorkstationId Setter
// 工作站id
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetWorkstationId(workstationId string) error {
    r.workstationId = workstationId
    r.Set("workstation_id", workstationId)
    return nil
}

// WorkstationId Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetWorkstationId() string {
    return r.workstationId
}
// WorkstationName Setter
// 工作站名称
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetWorkstationName(workstationName string) error {
    r.workstationName = workstationName
    r.Set("workstation_name", workstationName)
    return nil
}

// WorkstationName Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetWorkstationName() string {
    return r.workstationName
}
// PaymentList Setter
// 支付方式
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetPaymentList(paymentList string) error {
    r.paymentList = paymentList
    r.Set("payment_list", paymentList)
    return nil
}

// PaymentList Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetPaymentList() string {
    return r.paymentList
}
// PromotionList Setter
// 优惠列表
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetPromotionList(promotionList string) error {
    r.promotionList = promotionList
    r.Set("promotion_list", promotionList)
    return nil
}

// PromotionList Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetPromotionList() string {
    return r.promotionList
}
// GoodsList Setter
// 卖品列表
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetGoodsList(goodsList string) error {
    r.goodsList = goodsList
    r.Set("goods_list", goodsList)
    return nil
}

// GoodsList Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetGoodsList() string {
    return r.goodsList
}
// Mobile Setter
// 手机号
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetMobile() string {
    return r.mobile
}
// OperatorUserId Setter
// 管理员
func (r *TaobaoLarkIotOrderConfirmorderRequest) SetOperatorUserId(operatorUserId string) error {
    r.operatorUserId = operatorUserId
    r.Set("operator_user_id", operatorUserId)
    return nil
}

// OperatorUserId Getter
func (r TaobaoLarkIotOrderConfirmorderRequest) GetOperatorUserId() string {
    return r.operatorUserId
}