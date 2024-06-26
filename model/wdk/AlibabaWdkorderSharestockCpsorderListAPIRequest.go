package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockCpsorderListAPIRequest cps正向分销订单批量回流 API请求
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
type AlibabaWdkorderSharestockCpsorderListAPIRequest struct {
	model.Params
	// 入参
	_cpsOrderRequest *CpsOrderRequest
}

// NewAlibabaWdkorderSharestockCpsorderListRequest 初始化AlibabaWdkorderSharestockCpsorderListAPIRequest对象
func NewAlibabaWdkorderSharestockCpsorderListRequest() *AlibabaWdkorderSharestockCpsorderListAPIRequest {
	return &AlibabaWdkorderSharestockCpsorderListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkorderSharestockCpsorderListAPIRequest) Reset() {
	r._cpsOrderRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.cpsorder.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpsOrderRequest is CpsOrderRequest Setter
// 入参
func (r *AlibabaWdkorderSharestockCpsorderListAPIRequest) SetCpsOrderRequest(_cpsOrderRequest *CpsOrderRequest) error {
	r._cpsOrderRequest = _cpsOrderRequest
	r.Set("cps_order_request", _cpsOrderRequest)
	return nil
}

// GetCpsOrderRequest CpsOrderRequest Getter
func (r AlibabaWdkorderSharestockCpsorderListAPIRequest) GetCpsOrderRequest() *CpsOrderRequest {
	return r._cpsOrderRequest
}

var poolAlibabaWdkorderSharestockCpsorderListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkorderSharestockCpsorderListRequest()
	},
}

// GetAlibabaWdkorderSharestockCpsorderListRequest 从 sync.Pool 获取 AlibabaWdkorderSharestockCpsorderListAPIRequest
func GetAlibabaWdkorderSharestockCpsorderListAPIRequest() *AlibabaWdkorderSharestockCpsorderListAPIRequest {
	return poolAlibabaWdkorderSharestockCpsorderListAPIRequest.Get().(*AlibabaWdkorderSharestockCpsorderListAPIRequest)
}

// ReleaseAlibabaWdkorderSharestockCpsorderListAPIRequest 将 AlibabaWdkorderSharestockCpsorderListAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkorderSharestockCpsorderListAPIRequest(v *AlibabaWdkorderSharestockCpsorderListAPIRequest) {
	v.Reset()
	poolAlibabaWdkorderSharestockCpsorderListAPIRequest.Put(v)
}
