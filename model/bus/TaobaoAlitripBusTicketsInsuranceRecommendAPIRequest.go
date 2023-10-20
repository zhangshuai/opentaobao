package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripbusticketsinsurancerecommendAPIRequest 汽车票保险推荐 API请求
// taobao.alitrip.bus.tickets.insurance.recommend
//
// 获取推荐保险内容
type TaobaoalitripbusticketsinsurancerecommendAPIRequest struct {
	model.Params
	// 请求对象
	_recommendReq *TopStandardInsRecommendRequest
}

// NewTaobaoalitripbusticketsinsurancerecommendRequest 初始化TaobaoalitripbusticketsinsurancerecommendAPIRequest对象
func NewTaobaoalitripbusticketsinsurancerecommendRequest() *TaobaoalitripbusticketsinsurancerecommendAPIRequest {
	return &TaobaoalitripbusticketsinsurancerecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripbusticketsinsurancerecommendAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.bus.tickets.insurance.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripbusticketsinsurancerecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripbusticketsinsurancerecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecommendReq is RecommendReq Setter
// 请求对象
func (r *TaobaoalitripbusticketsinsurancerecommendAPIRequest) SetRecommendReq(_recommendReq *TopStandardInsRecommendRequest) error {
	r._recommendReq = _recommendReq
	r.Set("recommend_req", _recommendReq)
	return nil
}

// GetRecommendReq RecommendReq Getter
func (r TaobaoalitripbusticketsinsurancerecommendAPIRequest) GetRecommendReq() *TopStandardInsRecommendRequest {
	return r._recommendReq
}
