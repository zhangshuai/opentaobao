package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreRefusedAPIRequest Pos端门店拒单 API请求
// taobao.omniorder.store.refused
//
// ISV Pos端门店拒单，通知星盘
type TaobaoOmniorderStoreRefusedAPIRequest struct {
	model.Params
	// 子订单列表
	_subOrderList []SubOrder
	// 跟踪Id
	_traceId string
	// 淘宝交易主订单ID
	_tid int64
	// ISV的系统时间
	_reportTimestamp int64
}

// NewTaobaoOmniorderStoreRefusedRequest 初始化TaobaoOmniorderStoreRefusedAPIRequest对象
func NewTaobaoOmniorderStoreRefusedRequest() *TaobaoOmniorderStoreRefusedAPIRequest {
	return &TaobaoOmniorderStoreRefusedAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreRefusedAPIRequest) Reset() {
	r._subOrderList = r._subOrderList[:0]
	r._traceId = ""
	r._tid = 0
	r._reportTimestamp = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.refused"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderList is SubOrderList Setter
// 子订单列表
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetSubOrderList(_subOrderList []SubOrder) error {
	r._subOrderList = _subOrderList
	r.Set("sub_order_list", _subOrderList)
	return nil
}

// GetSubOrderList SubOrderList Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetSubOrderList() []SubOrder {
	return r._subOrderList
}

// SetTraceId is TraceId Setter
// 跟踪Id
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetTid is Tid Setter
// 淘宝交易主订单ID
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReportTimestamp is ReportTimestamp Setter
// ISV的系统时间
func (r *TaobaoOmniorderStoreRefusedAPIRequest) SetReportTimestamp(_reportTimestamp int64) error {
	r._reportTimestamp = _reportTimestamp
	r.Set("report_timestamp", _reportTimestamp)
	return nil
}

// GetReportTimestamp ReportTimestamp Getter
func (r TaobaoOmniorderStoreRefusedAPIRequest) GetReportTimestamp() int64 {
	return r._reportTimestamp
}

var poolTaobaoOmniorderStoreRefusedAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreRefusedRequest()
	},
}

// GetTaobaoOmniorderStoreRefusedRequest 从 sync.Pool 获取 TaobaoOmniorderStoreRefusedAPIRequest
func GetTaobaoOmniorderStoreRefusedAPIRequest() *TaobaoOmniorderStoreRefusedAPIRequest {
	return poolTaobaoOmniorderStoreRefusedAPIRequest.Get().(*TaobaoOmniorderStoreRefusedAPIRequest)
}

// ReleaseTaobaoOmniorderStoreRefusedAPIRequest 将 TaobaoOmniorderStoreRefusedAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreRefusedAPIRequest(v *TaobaoOmniorderStoreRefusedAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreRefusedAPIRequest.Put(v)
}
