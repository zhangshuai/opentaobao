package iot

import (
	"sync"
)

// BusinessRecipeStepActionOpenDto 结构体
type BusinessRecipeStepActionOpenDto struct {
	// 指令名
	ActionName string `json:"action_name,omitempty" xml:"action_name,omitempty"`
	// 指令值
	ActionValue string `json:"action_value,omitempty" xml:"action_value,omitempty"`
	// 指令类型：0:整数型,1:枚举类型,2:浮点类型,3:布尔型,4:字符串,5:时间型,6:JSON对象
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 指令顺序号
	Sequence int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
}

var poolBusinessRecipeStepActionOpenDto = sync.Pool{
	New: func() any {
		return new(BusinessRecipeStepActionOpenDto)
	},
}

// GetBusinessRecipeStepActionOpenDto() 从对象池中获取BusinessRecipeStepActionOpenDto
func GetBusinessRecipeStepActionOpenDto() *BusinessRecipeStepActionOpenDto {
	return poolBusinessRecipeStepActionOpenDto.Get().(*BusinessRecipeStepActionOpenDto)
}

// ReleaseBusinessRecipeStepActionOpenDto 释放BusinessRecipeStepActionOpenDto
func ReleaseBusinessRecipeStepActionOpenDto(v *BusinessRecipeStepActionOpenDto) {
	v.ActionName = ""
	v.ActionValue = ""
	v.ActionType = 0
	v.Sequence = 0
	poolBusinessRecipeStepActionOpenDto.Put(v)
}
