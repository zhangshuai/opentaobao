package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+平台取消订单 API请求
alibaba.nlife.b2c.trade.cancel

零售+平台取消订单接口
*/
type AlibabaNlifeB2cTradeCancelRequest struct {
    model.Params
    // 零售+平台订单号，和out_trade_no不能同时为空
    tradeNo   string
    // 外部订单号，和trade_no不能同时为空
    outTradeNo   string
    // 零售+门店号
    storeId   string
}

// 初始化AlibabaNlifeB2cTradeCancelRequest对象
func NewAlibabaNlifeB2cTradeCancelRequest() *AlibabaNlifeB2cTradeCancelRequest{
    return &AlibabaNlifeB2cTradeCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeCancelRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 零售+平台订单号，和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeCancelRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeB2cTradeCancelRequest) GetTradeNo() string {
    return r.tradeNo
}
// OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeCancelRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r AlibabaNlifeB2cTradeCancelRequest) GetOutTradeNo() string {
    return r.outTradeNo
}
// StoreId Setter
// 零售+门店号
func (r *AlibabaNlifeB2cTradeCancelRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradeCancelRequest) GetStoreId() string {
    return r.storeId
}