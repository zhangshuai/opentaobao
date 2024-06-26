package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest 前置校验商品是否可下单作业 API请求
// alibaba.tcls.aelophy.merchant.channel.order.precheck
//
// 前置校验商品是否可下单作业
type AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest struct {
	model.Params
	// 前置校验入参
	_preCheckRequest *PreCheckRequest
}

// NewAlibabaTclsAelophyMerchantChannelOrderPrecheckRequest 初始化AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderPrecheckRequest() *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) Reset() {
	r._preCheckRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.precheck"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPreCheckRequest is PreCheckRequest Setter
// 前置校验入参
func (r *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) SetPreCheckRequest(_preCheckRequest *PreCheckRequest) error {
	r._preCheckRequest = _preCheckRequest
	r.Set("pre_check_request", _preCheckRequest)
	return nil
}

// GetPreCheckRequest PreCheckRequest Getter
func (r AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) GetPreCheckRequest() *PreCheckRequest {
	return r._preCheckRequest
}

var poolAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelOrderPrecheckRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderPrecheckRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest
func GetAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest() *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest 将 AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest(v *AlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderPrecheckAPIRequest.Put(v)
}
