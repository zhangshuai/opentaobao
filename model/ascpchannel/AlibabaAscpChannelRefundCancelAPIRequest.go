package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundCancelAPIRequest 渠道退款单撤销 API请求
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
type AlibabaAscpChannelRefundCancelAPIRequest struct {
	model.Params
	// 入参
	_cancelRefundOrderRequest *Cancelrefundorderrequest
}

// NewAlibabaAscpChannelRefundCancelRequest 初始化AlibabaAscpChannelRefundCancelAPIRequest对象
func NewAlibabaAscpChannelRefundCancelRequest() *AlibabaAscpChannelRefundCancelAPIRequest {
	return &AlibabaAscpChannelRefundCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpChannelRefundCancelAPIRequest) Reset() {
	r._cancelRefundOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.channel.refund.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRefundOrderRequest is CancelRefundOrderRequest Setter
// 入参
func (r *AlibabaAscpChannelRefundCancelAPIRequest) SetCancelRefundOrderRequest(_cancelRefundOrderRequest *Cancelrefundorderrequest) error {
	r._cancelRefundOrderRequest = _cancelRefundOrderRequest
	r.Set("cancel_refund_order_request", _cancelRefundOrderRequest)
	return nil
}

// GetCancelRefundOrderRequest CancelRefundOrderRequest Getter
func (r AlibabaAscpChannelRefundCancelAPIRequest) GetCancelRefundOrderRequest() *Cancelrefundorderrequest {
	return r._cancelRefundOrderRequest
}

var poolAlibabaAscpChannelRefundCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpChannelRefundCancelRequest()
	},
}

// GetAlibabaAscpChannelRefundCancelRequest 从 sync.Pool 获取 AlibabaAscpChannelRefundCancelAPIRequest
func GetAlibabaAscpChannelRefundCancelAPIRequest() *AlibabaAscpChannelRefundCancelAPIRequest {
	return poolAlibabaAscpChannelRefundCancelAPIRequest.Get().(*AlibabaAscpChannelRefundCancelAPIRequest)
}

// ReleaseAlibabaAscpChannelRefundCancelAPIRequest 将 AlibabaAscpChannelRefundCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpChannelRefundCancelAPIRequest(v *AlibabaAscpChannelRefundCancelAPIRequest) {
	v.Reset()
	poolAlibabaAscpChannelRefundCancelAPIRequest.Put(v)
}
