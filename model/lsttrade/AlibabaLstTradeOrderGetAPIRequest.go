package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通交易订单查询--品牌商视角 API请求
alibaba.lst.trade.order.get

根据订单id查询零售通交易订单
*/
type AlibabaLstTradeOrderGetAPIRequest struct {
    model.Params
    // 主订单id
    _mainOrderId   int64
    // 子订单id
    _subOrderId   int64
}

// 初始化AlibabaLstTradeOrderGetAPIRequest对象
func NewAlibabaLstTradeOrderGetRequest() *AlibabaLstTradeOrderGetAPIRequest{
    return &AlibabaLstTradeOrderGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeOrderGetAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeOrderGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 主订单id
func (r *AlibabaLstTradeOrderGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r AlibabaLstTradeOrderGetAPIRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// SubOrderId Setter
// 子订单id
func (r *AlibabaLstTradeOrderGetAPIRequest) SetSubOrderId(_subOrderId int64) error {
    r._subOrderId = _subOrderId
    r.Set("sub_order_id", _subOrderId)
    return nil
}

// SubOrderId Getter
func (r AlibabaLstTradeOrderGetAPIRequest) GetSubOrderId() int64 {
    return r._subOrderId
}