package alsc

import (
	"sync"
)

// PointRuleOpenInfo 结构体
type PointRuleOpenInfo struct {
	// 创建者
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 规则名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 规则的业务ID
	RuleId string `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	// 更新者
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 更新者姓名
	UpdateByName string `json:"update_by_name,omitempty" xml:"update_by_name,omitempty"`
	// 创建者姓名
	CreateByName string `json:"create_by_name,omitempty" xml:"create_by_name,omitempty"`
	// 扩展信息
	ExtInfo *Extinfo `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 积分获取规则
	PointAdditionRule *PointAdditionRuleOpenInfo `json:"point_addition_rule,omitempty" xml:"point_addition_rule,omitempty"`
	// 积分清零规则
	PointClearRule *PointClearRuleOpenInfo `json:"point_clear_rule,omitempty" xml:"point_clear_rule,omitempty"`
	// 积分扣减规则
	PointDeductionRule *PointDeductionRuleOpenInfo `json:"point_deduction_rule,omitempty" xml:"point_deduction_rule,omitempty"`
	// 逻辑删除标志
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolPointRuleOpenInfo = sync.Pool{
	New: func() any {
		return new(PointRuleOpenInfo)
	},
}

// GetPointRuleOpenInfo() 从对象池中获取PointRuleOpenInfo
func GetPointRuleOpenInfo() *PointRuleOpenInfo {
	return poolPointRuleOpenInfo.Get().(*PointRuleOpenInfo)
}

// ReleasePointRuleOpenInfo 释放PointRuleOpenInfo
func ReleasePointRuleOpenInfo(v *PointRuleOpenInfo) {
	v.CreateBy = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Name = ""
	v.RuleId = ""
	v.UpdateBy = ""
	v.UpdateByName = ""
	v.CreateByName = ""
	v.ExtInfo = nil
	v.PointAdditionRule = nil
	v.PointClearRule = nil
	v.PointDeductionRule = nil
	v.Deleted = false
	poolPointRuleOpenInfo.Put(v)
}
