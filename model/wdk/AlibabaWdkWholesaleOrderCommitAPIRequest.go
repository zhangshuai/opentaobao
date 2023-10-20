package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkwholesaleordercommitAPIRequest 盒马帮采购确认订单接口 API请求
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
type AlibabawdkwholesaleordercommitAPIRequest struct {
	model.Params
	// 采购单信息
	_orderCommitReq *OrderCommitReq
}

// NewAlibabawdkwholesaleordercommitRequest 初始化AlibabawdkwholesaleordercommitAPIRequest对象
func NewAlibabawdkwholesaleordercommitRequest() *AlibabawdkwholesaleordercommitAPIRequest {
	return &AlibabawdkwholesaleordercommitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkwholesaleordercommitAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.wholesale.order.commit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkwholesaleordercommitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkwholesaleordercommitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCommitReq is OrderCommitReq Setter
// 采购单信息
func (r *AlibabawdkwholesaleordercommitAPIRequest) SetOrderCommitReq(_orderCommitReq *OrderCommitReq) error {
	r._orderCommitReq = _orderCommitReq
	r.Set("order_commit_req", _orderCommitReq)
	return nil
}

// GetOrderCommitReq OrderCommitReq Getter
func (r AlibabawdkwholesaleordercommitAPIRequest) GetOrderCommitReq() *OrderCommitReq {
	return r._orderCommitReq
}
