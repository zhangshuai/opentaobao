package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单延住接口 API请求
taobao.xhotel.order.alipayface.extend

信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新
*/
type TaobaoXhotelOrderAlipayfaceExtendRequest struct {
    model.Params
    // 延住后的离店日期(最多总共入住天数不能超过9间夜)
    checkOut   string
    // 阿里旅行订单id,和outOrderId必须至少传入一个
    tid   int64
    // 卖家系统订单号,和tid必须至少传入一个
    outOrderId   string
    // 延住房费,注意不包含原有的房费金额,单位为分
    extendFee   int64
    // 延住日期段的每日房价信息,注意不包括原有的日期房价
    extendDailyPriceInfo   string
}

// 初始化TaobaoXhotelOrderAlipayfaceExtendRequest对象
func NewTaobaoXhotelOrderAlipayfaceExtendRequest() *TaobaoXhotelOrderAlipayfaceExtendRequest{
    return &TaobaoXhotelOrderAlipayfaceExtendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.extend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CheckOut Setter
// 延住后的离店日期(最多总共入住天数不能超过9间夜)
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetCheckOut() string {
    return r.checkOut
}
// Tid Setter
// 阿里旅行订单id,和outOrderId必须至少传入一个
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetTid() int64 {
    return r.tid
}
// OutOrderId Setter
// 卖家系统订单号,和tid必须至少传入一个
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetOutOrderId() string {
    return r.outOrderId
}
// ExtendFee Setter
// 延住房费,注意不包含原有的房费金额,单位为分
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetExtendFee(extendFee int64) error {
    r.extendFee = extendFee
    r.Set("extend_fee", extendFee)
    return nil
}

// ExtendFee Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetExtendFee() int64 {
    return r.extendFee
}
// ExtendDailyPriceInfo Setter
// 延住日期段的每日房价信息,注意不包括原有的日期房价
func (r *TaobaoXhotelOrderAlipayfaceExtendRequest) SetExtendDailyPriceInfo(extendDailyPriceInfo string) error {
    r.extendDailyPriceInfo = extendDailyPriceInfo
    r.Set("extend_daily_price_info", extendDailyPriceInfo)
    return nil
}

// ExtendDailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceExtendRequest) GetExtendDailyPriceInfo() string {
    return r.extendDailyPriceInfo
}