package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterSubscribeQueryAPIRequest 天猫服务订购信息查询接口 API请求
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
type TmallMallitemcenterSubscribeQueryAPIRequest struct {
	model.Params
	// 入参query
	_query *Spb2bOderQuery
}

// NewTmallMallitemcenterSubscribeQueryRequest 初始化TmallMallitemcenterSubscribeQueryAPIRequest对象
func NewTmallMallitemcenterSubscribeQueryRequest() *TmallMallitemcenterSubscribeQueryAPIRequest {
	return &TmallMallitemcenterSubscribeQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMallitemcenterSubscribeQueryAPIRequest) Reset() {
	r._query = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetApiMethodName() string {
	return "tmall.mallitemcenter.subscribe.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 入参query
func (r *TmallMallitemcenterSubscribeQueryAPIRequest) SetQuery(_query *Spb2bOderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TmallMallitemcenterSubscribeQueryAPIRequest) GetQuery() *Spb2bOderQuery {
	return r._query
}

var poolTmallMallitemcenterSubscribeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMallitemcenterSubscribeQueryRequest()
	},
}

// GetTmallMallitemcenterSubscribeQueryRequest 从 sync.Pool 获取 TmallMallitemcenterSubscribeQueryAPIRequest
func GetTmallMallitemcenterSubscribeQueryAPIRequest() *TmallMallitemcenterSubscribeQueryAPIRequest {
	return poolTmallMallitemcenterSubscribeQueryAPIRequest.Get().(*TmallMallitemcenterSubscribeQueryAPIRequest)
}

// ReleaseTmallMallitemcenterSubscribeQueryAPIRequest 将 TmallMallitemcenterSubscribeQueryAPIRequest 放入 sync.Pool
func ReleaseTmallMallitemcenterSubscribeQueryAPIRequest(v *TmallMallitemcenterSubscribeQueryAPIRequest) {
	v.Reset()
	poolTmallMallitemcenterSubscribeQueryAPIRequest.Put(v)
}
