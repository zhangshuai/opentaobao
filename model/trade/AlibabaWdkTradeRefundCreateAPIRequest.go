package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundCreateAPIRequest 外部渠道逆向订单创建 API请求
// alibaba.wdk.trade.refund.create
//
// 该接口是创建退货订单的服务。当外部渠道发起退款后，调用此接口可以完成五道口底层交易、履约、配送等一系列流程进行退货。
type AlibabaWdkTradeRefundCreateAPIRequest struct {
	model.Params
	// 退货请求
	_refundGoodsCreateRequest *RefundGoodsCreateRequest
}

// NewAlibabaWdkTradeRefundCreateRequest 初始化AlibabaWdkTradeRefundCreateAPIRequest对象
func NewAlibabaWdkTradeRefundCreateRequest() *AlibabaWdkTradeRefundCreateAPIRequest {
	return &AlibabaWdkTradeRefundCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeRefundCreateAPIRequest) Reset() {
	r._refundGoodsCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeRefundCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.refund.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeRefundCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeRefundCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundGoodsCreateRequest is RefundGoodsCreateRequest Setter
// 退货请求
func (r *AlibabaWdkTradeRefundCreateAPIRequest) SetRefundGoodsCreateRequest(_refundGoodsCreateRequest *RefundGoodsCreateRequest) error {
	r._refundGoodsCreateRequest = _refundGoodsCreateRequest
	r.Set("refund_goods_create_request", _refundGoodsCreateRequest)
	return nil
}

// GetRefundGoodsCreateRequest RefundGoodsCreateRequest Getter
func (r AlibabaWdkTradeRefundCreateAPIRequest) GetRefundGoodsCreateRequest() *RefundGoodsCreateRequest {
	return r._refundGoodsCreateRequest
}

var poolAlibabaWdkTradeRefundCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeRefundCreateRequest()
	},
}

// GetAlibabaWdkTradeRefundCreateRequest 从 sync.Pool 获取 AlibabaWdkTradeRefundCreateAPIRequest
func GetAlibabaWdkTradeRefundCreateAPIRequest() *AlibabaWdkTradeRefundCreateAPIRequest {
	return poolAlibabaWdkTradeRefundCreateAPIRequest.Get().(*AlibabaWdkTradeRefundCreateAPIRequest)
}

// ReleaseAlibabaWdkTradeRefundCreateAPIRequest 将 AlibabaWdkTradeRefundCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeRefundCreateAPIRequest(v *AlibabaWdkTradeRefundCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeRefundCreateAPIRequest.Put(v)
}
