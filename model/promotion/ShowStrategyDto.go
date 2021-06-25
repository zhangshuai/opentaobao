package promotion

// ShowStrategyDto 
type ShowStrategyDto struct {

    // 投放模式
    Mode   string `json:"mode,omitempty"`

    // 投放计划code
    Code   string `json:"code,omitempty"`

    // 是否投放计划所有规则通过
    AllRulePassed   bool `json:"all_rule_passed,omitempty"`

    // 投放计划安全码
    Asac   string `json:"asac,omitempty"`

    // 算法容灾结果
    AlgorithmFailover   bool `json:"algorithm_failover,omitempty"`

    // 投放计划规则
    ShowRules   []ShowRuleDto `json:"show_rules,omitempty"`

}