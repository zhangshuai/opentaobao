package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店接单接口 API请求
taobao.omniorder.store.accpeted

ISV Pos端门店接单，通知星盘
*/
type TaobaoOmniorderStoreAccpetedAPIRequest struct {
    model.Params
    // 淘宝交易主订单ID
    _tid   int64
    // 子订单列表
    _subOrderList   []StoreAcceptedResult
    // ISV系统上报时间
    _reportTimestamp   int64
    // 跟踪Id
    _traceId   string
}

// 初始化TaobaoOmniorderStoreAccpetedAPIRequest对象
func NewTaobaoOmniorderStoreAccpetedRequest() *TaobaoOmniorderStoreAccpetedAPIRequest{
    return &TaobaoOmniorderStoreAccpetedAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.accpeted"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetTid() int64 {
    return r._tid
}
// SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetSubOrderList(_subOrderList []StoreAcceptedResult) error {
    r._subOrderList = _subOrderList
    r.Set("sub_order_list", _subOrderList)
    return nil
}

// SubOrderList Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetSubOrderList() []StoreAcceptedResult {
    return r._subOrderList
}
// ReportTimestamp Setter
// ISV系统上报时间
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
    r._reportTimestamp = _reportTimestamp
    r.Set("report_timestamp", _reportTimestamp)
    return nil
}

// ReportTimestamp Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetReportTimestamp() int64 {
    return r._reportTimestamp
}
// TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreAccpetedAPIRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r TaobaoOmniorderStoreAccpetedAPIRequest) GetTraceId() string {
    return r._traceId
}