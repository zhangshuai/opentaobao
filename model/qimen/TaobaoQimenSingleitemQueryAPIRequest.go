package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSingleitemQueryAPIRequest 商品查询接口 API请求
// taobao.qimen.singleitem.query
//
// 商品查询接口
type TaobaoQimenSingleitemQueryAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// NewTaobaoQimenSingleitemQueryRequest 初始化TaobaoQimenSingleitemQueryAPIRequest对象
func NewTaobaoQimenSingleitemQueryRequest() *TaobaoQimenSingleitemQueryAPIRequest {
	return &TaobaoQimenSingleitemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenSingleitemQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.singleitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenSingleitemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenSingleitemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenSingleitemQueryAPIRequest) SetRequest(_request *RequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenSingleitemQueryAPIRequest) GetRequest() *RequestDo {
	return r._request
}
