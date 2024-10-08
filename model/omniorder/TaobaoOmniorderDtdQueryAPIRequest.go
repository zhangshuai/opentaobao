package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdQueryAPIRequest 门店自送根据核销码查订单 API请求
// taobao.omniorder.dtd.query
//
// 门店自送根据核销码码查询订单信息
type TaobaoOmniorderDtdQueryAPIRequest struct {
	model.Params
	// 核销码
	_code string
}

// NewTaobaoOmniorderDtdQueryRequest 初始化TaobaoOmniorderDtdQueryAPIRequest对象
func NewTaobaoOmniorderDtdQueryRequest() *TaobaoOmniorderDtdQueryAPIRequest {
	return &TaobaoOmniorderDtdQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderDtdQueryAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdQueryAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.dtd.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderDtdQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 核销码
func (r *TaobaoOmniorderDtdQueryAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaoOmniorderDtdQueryAPIRequest) GetCode() string {
	return r._code
}

var poolTaobaoOmniorderDtdQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderDtdQueryRequest()
	},
}

// GetTaobaoOmniorderDtdQueryRequest 从 sync.Pool 获取 TaobaoOmniorderDtdQueryAPIRequest
func GetTaobaoOmniorderDtdQueryAPIRequest() *TaobaoOmniorderDtdQueryAPIRequest {
	return poolTaobaoOmniorderDtdQueryAPIRequest.Get().(*TaobaoOmniorderDtdQueryAPIRequest)
}

// ReleaseTaobaoOmniorderDtdQueryAPIRequest 将 TaobaoOmniorderDtdQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderDtdQueryAPIRequest(v *TaobaoOmniorderDtdQueryAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderDtdQueryAPIRequest.Put(v)
}
