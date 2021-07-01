package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商获得用户身份信息 API请求
alibaba.aliqin.tcc.trade.identity.get

天猫网厅运营商官方旗舰店获取用户身份信息
*/
type AlibabaAliqinTccTradeIdentityGetAPIRequest struct {
    model.Params
    // 订单编号
    _bizOrderId   int64
    // 店铺名称
    _sellerNick   string
}

// 初始化AlibabaAliqinTccTradeIdentityGetAPIRequest对象
func NewAlibabaAliqinTccTradeIdentityGetRequest() *AlibabaAliqinTccTradeIdentityGetAPIRequest{
    return &AlibabaAliqinTccTradeIdentityGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinTccTradeIdentityGetAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.tcc.trade.identity.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinTccTradeIdentityGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizOrderId Setter
// 订单编号
func (r *AlibabaAliqinTccTradeIdentityGetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
    r._bizOrderId = _bizOrderId
    r.Set("biz_order_id", _bizOrderId)
    return nil
}

// BizOrderId Getter
func (r AlibabaAliqinTccTradeIdentityGetAPIRequest) GetBizOrderId() int64 {
    return r._bizOrderId
}
// SellerNick Setter
// 店铺名称
func (r *AlibabaAliqinTccTradeIdentityGetAPIRequest) SetSellerNick(_sellerNick string) error {
    r._sellerNick = _sellerNick
    r.Set("seller_nick", _sellerNick)
    return nil
}

// SellerNick Getter
func (r AlibabaAliqinTccTradeIdentityGetAPIRequest) GetSellerNick() string {
    return r._sellerNick
}