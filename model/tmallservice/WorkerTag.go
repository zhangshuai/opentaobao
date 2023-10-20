package tmallservice

import (
	"sync"
)

// WorkerTag 结构体
type WorkerTag struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 能力列表 key为能力，value则是对应能力的详细描述.  effective: 当前能力是否已经激活。  serviceCode: 工人具备的能力。  abilityLevel:  0: 没有等级，1: 初级工人  2: 中级工人  3:高级工人 4:特级工人
	AbilityMap string `json:"ability_map,omitempty" xml:"ability_map,omitempty"`
	// 工人能力
	ResultData *WorkerTag `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 黑名单状态
	BlackList bool `json:"black_list,omitempty" xml:"black_list,omitempty"`
	// 飞单状态
	Degradation bool `json:"degradation,omitempty" xml:"degradation,omitempty"`
	// vip工人
	Vip bool `json:"vip,omitempty" xml:"vip,omitempty"`
	// 大促工人
	BigPromotion bool `json:"big_promotion,omitempty" xml:"big_promotion,omitempty"`
}

var poolWorkerTag = sync.Pool{
	New: func() any {
		return new(WorkerTag)
	},
}

// GetWorkerTag() 从对象池中获取WorkerTag
func GetWorkerTag() *WorkerTag {
	return poolWorkerTag.Get().(*WorkerTag)
}

// ReleaseWorkerTag 释放WorkerTag
func ReleaseWorkerTag(v *WorkerTag) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.AbilityMap = ""
	v.ResultData = nil
	v.Success = false
	v.BlackList = false
	v.Degradation = false
	v.Vip = false
	v.BigPromotion = false
	poolWorkerTag.Put(v)
}
