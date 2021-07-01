package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易号获取物流宝订单 API请求
taobao.wlb.tradeorder.get

根据交易类型和交易id查询物流宝订单详情
*/
type TaobaoWlbTradeorderGetAPIRequest struct {
    model.Params
    // 子交易号
    _subTradeId   string
    // 指定交易类型的交易号
    _tradeId   string
    // 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
    _tradeType   string
}

// 初始化TaobaoWlbTradeorderGetAPIRequest对象
func NewTaobaoWlbTradeorderGetRequest() *TaobaoWlbTradeorderGetAPIRequest{
    return &TaobaoWlbTradeorderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbTradeorderGetAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.tradeorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbTradeorderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubTradeId Setter
// 子交易号
func (r *TaobaoWlbTradeorderGetAPIRequest) SetSubTradeId(_subTradeId string) error {
    r._subTradeId = _subTradeId
    r.Set("sub_trade_id", _subTradeId)
    return nil
}

// SubTradeId Getter
func (r TaobaoWlbTradeorderGetAPIRequest) GetSubTradeId() string {
    return r._subTradeId
}
// TradeId Setter
// 指定交易类型的交易号
func (r *TaobaoWlbTradeorderGetAPIRequest) SetTradeId(_tradeId string) error {
    r._tradeId = _tradeId
    r.Set("trade_id", _tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoWlbTradeorderGetAPIRequest) GetTradeId() string {
    return r._tradeId
}
// TradeType Setter
// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
func (r *TaobaoWlbTradeorderGetAPIRequest) SetTradeType(_tradeType string) error {
    r._tradeType = _tradeType
    r.Set("trade_type", _tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoWlbTradeorderGetAPIRequest) GetTradeType() string {
    return r._tradeType
}