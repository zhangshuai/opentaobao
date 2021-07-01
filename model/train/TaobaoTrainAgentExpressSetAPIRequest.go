package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息 API请求
taobao.train.agent.express.set

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetAPIRequest struct {
    model.Params
    // 订单号
    _mainOrderId   int64
    // 物流单号
    _expressId   string
    // 发货地址
    _addr   string
    // 手机号
    _mobile   string
    // 代理商id
    _agentId   int64
    // 物流公司:SF,EMS
    _expressName   string
}

// 初始化TaobaoTrainAgentExpressSetAPIRequest对象
func NewTaobaoTrainAgentExpressSetRequest() *TaobaoTrainAgentExpressSetAPIRequest{
    return &TaobaoTrainAgentExpressSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentExpressSetAPIRequest) GetApiMethodName() string {
    return "taobao.train.agent.express.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentExpressSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 订单号
func (r *TaobaoTrainAgentExpressSetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentExpressSetAPIRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// ExpressId Setter
// 物流单号
func (r *TaobaoTrainAgentExpressSetAPIRequest) SetExpressId(_expressId string) error {
    r._expressId = _expressId
    r.Set("express_id", _expressId)
    return nil
}

// ExpressId Getter
func (r TaobaoTrainAgentExpressSetAPIRequest) GetExpressId() string {
    return r._expressId
}
// Addr Setter
// 发货地址
func (r *TaobaoTrainAgentExpressSetAPIRequest) SetAddr(_addr string) error {
    r._addr = _addr
    r.Set("addr", _addr)
    return nil
}

// Addr Getter
func (r TaobaoTrainAgentExpressSetAPIRequest) GetAddr() string {
    return r._addr
}
// Mobile Setter
// 手机号
func (r *TaobaoTrainAgentExpressSetAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoTrainAgentExpressSetAPIRequest) GetMobile() string {
    return r._mobile
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentExpressSetAPIRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentExpressSetAPIRequest) GetAgentId() int64 {
    return r._agentId
}
// ExpressName Setter
// 物流公司:SF,EMS
func (r *TaobaoTrainAgentExpressSetAPIRequest) SetExpressName(_expressName string) error {
    r._expressName = _expressName
    r.Set("express_name", _expressName)
    return nil
}

// ExpressName Getter
func (r TaobaoTrainAgentExpressSetAPIRequest) GetExpressName() string {
    return r._expressName
}