package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmrulequerydishruleAPIRequest 查询品牌下的入会菜品规则 API请求
// alibaba.alsc.crm.rule.querydishrule
//
// 查询品牌下的入会菜品规则
type AlibabaalsccrmrulequerydishruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// NewAlibabaalsccrmrulequerydishruleRequest 初始化AlibabaalsccrmrulequerydishruleAPIRequest对象
func NewAlibabaalsccrmrulequerydishruleRequest() *AlibabaalsccrmrulequerydishruleAPIRequest {
	return &AlibabaalsccrmrulequerydishruleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmrulequerydishruleAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.rule.querydishrule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmrulequerydishruleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmrulequerydishruleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPlanRuleQueryOpenReq is ParamPlanRuleQueryOpenReq Setter
// 系统自动生成
func (r *AlibabaalsccrmrulequerydishruleAPIRequest) SetParamPlanRuleQueryOpenReq(_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq) error {
	r._paramPlanRuleQueryOpenReq = _paramPlanRuleQueryOpenReq
	r.Set("param_plan_rule_query_open_req", _paramPlanRuleQueryOpenReq)
	return nil
}

// GetParamPlanRuleQueryOpenReq ParamPlanRuleQueryOpenReq Getter
func (r AlibabaalsccrmrulequerydishruleAPIRequest) GetParamPlanRuleQueryOpenReq() *PlanRuleQueryOpenReq {
	return r._paramPlanRuleQueryOpenReq
}
