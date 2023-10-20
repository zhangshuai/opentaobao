package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenCombineitemSynchronizeAPIRequest 组合商品接口 API请求
// taobao.qimen.combineitem.synchronize
//
// ERP调用奇门的接口,将商品信息同步给WMS
type TaobaoQimenCombineitemSynchronizeAPIRequest struct {
	model.Params
	//
	_request *CombineItemSyncRequest
}

// NewTaobaoQimenCombineitemSynchronizeRequest 初始化TaobaoQimenCombineitemSynchronizeAPIRequest对象
func NewTaobaoQimenCombineitemSynchronizeRequest() *TaobaoQimenCombineitemSynchronizeAPIRequest {
	return &TaobaoQimenCombineitemSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.combineitem.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenCombineitemSynchronizeAPIRequest) SetRequest(_request *CombineItemSyncRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenCombineitemSynchronizeAPIRequest) GetRequest() *CombineItemSyncRequest {
	return r._request
}
