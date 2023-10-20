package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderQueryAPIRequest 订单详情查询 API请求
// taobao.train.agent.order.query
//
// 订单详情查询接口
type TaobaoTrainAgentOrderQueryAPIRequest struct {
	model.Params
	// rq
	_param0 *QueryOrderRq
}

// NewTaobaoTrainAgentOrderQueryRequest 初始化TaobaoTrainAgentOrderQueryAPIRequest对象
func NewTaobaoTrainAgentOrderQueryRequest() *TaobaoTrainAgentOrderQueryAPIRequest {
	return &TaobaoTrainAgentOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTrainAgentOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// rq
func (r *TaobaoTrainAgentOrderQueryAPIRequest) SetParam0(_param0 *QueryOrderRq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoTrainAgentOrderQueryAPIRequest) GetParam0() *QueryOrderRq {
	return r._param0
}
