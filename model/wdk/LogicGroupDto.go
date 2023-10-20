package wdk

import (
	"sync"
)

// LogicGroupDto 结构体
type LogicGroupDto struct {
	// 分组名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分组序号，换购场景需要两个逻辑分组，一个序号为1的普通逻辑分组和一个序号为2的换购逻辑分组
	Number int64 `json:"number,omitempty" xml:"number,omitempty"`
	// 分组类型，1: 普通分组， 2: 换购分组
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 换购规则
	ExchangeRule *ExchangeRuleDto `json:"exchange_rule,omitempty" xml:"exchange_rule,omitempty"`
	// 优惠门槛
	Condition *ConditionDto `json:"condition,omitempty" xml:"condition,omitempty"`
	// 优惠效果
	Action *ActionDto `json:"action,omitempty" xml:"action,omitempty"`
	// 换购分组标识
	Exchange bool `json:"exchange,omitempty" xml:"exchange,omitempty"`
	// 是否生效分组（多分组情况下，可能为false，例如商品池换购，普通逻辑分组为false，换购分组为true）
	EffectiveGroup bool `json:"effective_group,omitempty" xml:"effective_group,omitempty"`
}

var poolLogicGroupDto = sync.Pool{
	New: func() any {
		return new(LogicGroupDto)
	},
}

// GetLogicGroupDto() 从对象池中获取LogicGroupDto
func GetLogicGroupDto() *LogicGroupDto {
	return poolLogicGroupDto.Get().(*LogicGroupDto)
}

// ReleaseLogicGroupDto 释放LogicGroupDto
func ReleaseLogicGroupDto(v *LogicGroupDto) {
	v.Name = ""
	v.Number = 0
	v.Type = 0
	v.ExchangeRule = nil
	v.Condition = nil
	v.Action = nil
	v.Exchange = false
	v.EffectiveGroup = false
	poolLogicGroupDto.Put(v)
}
