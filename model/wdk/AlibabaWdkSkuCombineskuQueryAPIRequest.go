package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuQueryAPIRequest 组合商品查询接口 API请求
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
type AlibabaWdkSkuCombineskuQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param *SkuQueryDo
}

// NewAlibabaWdkSkuCombineskuQueryRequest 初始化AlibabaWdkSkuCombineskuQueryAPIRequest对象
func NewAlibabaWdkSkuCombineskuQueryRequest() *AlibabaWdkSkuCombineskuQueryAPIRequest {
	return &AlibabaWdkSkuCombineskuQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuCombineskuQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuQueryAPIRequest) SetParam(_param *SkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuCombineskuQueryAPIRequest) GetParam() *SkuQueryDo {
	return r._param
}

var poolAlibabaWdkSkuCombineskuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuCombineskuQueryRequest()
	},
}

// GetAlibabaWdkSkuCombineskuQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuCombineskuQueryAPIRequest
func GetAlibabaWdkSkuCombineskuQueryAPIRequest() *AlibabaWdkSkuCombineskuQueryAPIRequest {
	return poolAlibabaWdkSkuCombineskuQueryAPIRequest.Get().(*AlibabaWdkSkuCombineskuQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuCombineskuQueryAPIRequest 将 AlibabaWdkSkuCombineskuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuCombineskuQueryAPIRequest(v *AlibabaWdkSkuCombineskuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuCombineskuQueryAPIRequest.Put(v)
}
