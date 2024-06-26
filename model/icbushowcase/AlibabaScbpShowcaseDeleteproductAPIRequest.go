package icbushowcase

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseDeleteproductAPIRequest 批量删除橱窗商品 API请求
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
type AlibabaScbpShowcaseDeleteproductAPIRequest struct {
	model.Params
	// 橱窗idList
	_windowIdList []string
}

// NewAlibabaScbpShowcaseDeleteproductRequest 初始化AlibabaScbpShowcaseDeleteproductAPIRequest对象
func NewAlibabaScbpShowcaseDeleteproductRequest() *AlibabaScbpShowcaseDeleteproductAPIRequest {
	return &AlibabaScbpShowcaseDeleteproductAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpShowcaseDeleteproductAPIRequest) Reset() {
	r._windowIdList = r._windowIdList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpShowcaseDeleteproductAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.showcase.deleteproduct"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpShowcaseDeleteproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpShowcaseDeleteproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWindowIdList is WindowIdList Setter
// 橱窗idList
func (r *AlibabaScbpShowcaseDeleteproductAPIRequest) SetWindowIdList(_windowIdList []string) error {
	r._windowIdList = _windowIdList
	r.Set("window_id_list", _windowIdList)
	return nil
}

// GetWindowIdList WindowIdList Getter
func (r AlibabaScbpShowcaseDeleteproductAPIRequest) GetWindowIdList() []string {
	return r._windowIdList
}

var poolAlibabaScbpShowcaseDeleteproductAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpShowcaseDeleteproductRequest()
	},
}

// GetAlibabaScbpShowcaseDeleteproductRequest 从 sync.Pool 获取 AlibabaScbpShowcaseDeleteproductAPIRequest
func GetAlibabaScbpShowcaseDeleteproductAPIRequest() *AlibabaScbpShowcaseDeleteproductAPIRequest {
	return poolAlibabaScbpShowcaseDeleteproductAPIRequest.Get().(*AlibabaScbpShowcaseDeleteproductAPIRequest)
}

// ReleaseAlibabaScbpShowcaseDeleteproductAPIRequest 将 AlibabaScbpShowcaseDeleteproductAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpShowcaseDeleteproductAPIRequest(v *AlibabaScbpShowcaseDeleteproductAPIRequest) {
	v.Reset()
	poolAlibabaScbpShowcaseDeleteproductAPIRequest.Put(v)
}
