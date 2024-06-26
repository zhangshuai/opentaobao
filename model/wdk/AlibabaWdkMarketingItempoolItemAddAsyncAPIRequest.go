package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest 商品池新增商品 API请求
// alibaba.wdk.marketing.itempool.item.add.async
//
// 新分组模型下新增商品
type AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest struct {
	model.Params
	// 阶梯商品信息
	_param0 []ItemPoolSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabaWdkMarketingItempoolItemAddAsyncRequest 初始化AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest对象
func NewAlibabaWdkMarketingItempoolItemAddAsyncRequest() *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest {
	return &AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) Reset() {
	r._param0 = r._param0[:0]
	r._param1 = nil
	r._version = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 阶梯商品信息
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) SetParam0(_param0 []ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetParam0() []ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) GetVersion() int64 {
	return r._version
}

var poolAlibabaWdkMarketingItempoolItemAddAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolItemAddAsyncRequest()
	},
}

// GetAlibabaWdkMarketingItempoolItemAddAsyncRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest
func GetAlibabaWdkMarketingItempoolItemAddAsyncAPIRequest() *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest {
	return poolAlibabaWdkMarketingItempoolItemAddAsyncAPIRequest.Get().(*AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolItemAddAsyncAPIRequest 将 AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolItemAddAsyncAPIRequest(v *AlibabaWdkMarketingItempoolItemAddAsyncAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolItemAddAsyncAPIRequest.Put(v)
}
