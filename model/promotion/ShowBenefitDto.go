package promotion

import (
	"sync"
)

// ShowBenefitDto 结构体
type ShowBenefitDto struct {
	// 权益规则列表
	ShowRules []ShowRuleDto `json:"show_rules,omitempty" xml:"show_rules>show_rule_dto,omitempty"`
	// 权益维度核销数据
	ShowBenefitInstances []ShowBenefitInstanceDto `json:"show_benefit_instances,omitempty" xml:"show_benefit_instances>show_benefit_instance_dto,omitempty"`
	// 权益code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 权益类型
	TypeDesc string `json:"type_desc,omitempty" xml:"type_desc,omitempty"`
	// 权益展示面额单位
	DisplayAmountUnit string `json:"display_amount_unit,omitempty" xml:"display_amount_unit,omitempty"`
	// 权益发放结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 权益标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 权益类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 权益扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 相对使用时间类型
	IntervalTimeUnit string `json:"interval_time_unit,omitempty" xml:"interval_time_unit,omitempty"`
	// 展示门槛
	DisplayStartFee string `json:"display_start_fee,omitempty" xml:"display_start_fee,omitempty"`
	// 权益发放类型
	SendMode string `json:"send_mode,omitempty" xml:"send_mode,omitempty"`
	// 权益生命周期
	SendLifeCycleState string `json:"send_life_cycle_state,omitempty" xml:"send_life_cycle_state,omitempty"`
	// 面额单位
	AmountUnit string `json:"amount_unit,omitempty" xml:"amount_unit,omitempty"`
	// 展示面额
	DisplayAmount string `json:"display_amount,omitempty" xml:"display_amount,omitempty"`
	// 使用时间类型
	EffectiveTimeMode string `json:"effective_time_mode,omitempty" xml:"effective_time_mode,omitempty"`
	// 素材
	Material string `json:"material,omitempty" xml:"material,omitempty"`
	// 加密动态面额参数
	EncryptedDynamicInfo string `json:"encrypted_dynamic_info,omitempty" xml:"encrypted_dynamic_info,omitempty"`
	// 权益安全码
	Asac string `json:"asac,omitempty" xml:"asac,omitempty"`
	// 发放开始时间
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 绝对使用时间开始
	EffectiveStart string `json:"effective_start,omitempty" xml:"effective_start,omitempty"`
	// 绝对使用时间结束
	EffectiveEnd string `json:"effective_end,omitempty" xml:"effective_end,omitempty"`
	// 权益面额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 相对使用时间长度
	EffectiveInterval int64 `json:"effective_interval,omitempty" xml:"effective_interval,omitempty"`
	// 门槛
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 权益是否可领
	CanWin bool `json:"can_win,omitempty" xml:"can_win,omitempty"`
	// 是否测试权益
	Test bool `json:"test,omitempty" xml:"test,omitempty"`
	// 是否已领
	HadWin bool `json:"had_win,omitempty" xml:"had_win,omitempty"`
	// 是否有库存
	HasInventory bool `json:"has_inventory,omitempty" xml:"has_inventory,omitempty"`
}

var poolShowBenefitDto = sync.Pool{
	New: func() any {
		return new(ShowBenefitDto)
	},
}

// GetShowBenefitDto() 从对象池中获取ShowBenefitDto
func GetShowBenefitDto() *ShowBenefitDto {
	return poolShowBenefitDto.Get().(*ShowBenefitDto)
}

// ReleaseShowBenefitDto 释放ShowBenefitDto
func ReleaseShowBenefitDto(v *ShowBenefitDto) {
	v.ShowRules = v.ShowRules[:0]
	v.ShowBenefitInstances = v.ShowBenefitInstances[:0]
	v.Code = ""
	v.TypeDesc = ""
	v.DisplayAmountUnit = ""
	v.EndDate = ""
	v.Title = ""
	v.Type = ""
	v.Feature = ""
	v.IntervalTimeUnit = ""
	v.DisplayStartFee = ""
	v.SendMode = ""
	v.SendLifeCycleState = ""
	v.AmountUnit = ""
	v.DisplayAmount = ""
	v.EffectiveTimeMode = ""
	v.Material = ""
	v.EncryptedDynamicInfo = ""
	v.Asac = ""
	v.StartDate = ""
	v.EffectiveStart = ""
	v.EffectiveEnd = ""
	v.Amount = 0
	v.EffectiveInterval = 0
	v.StartFee = 0
	v.CanWin = false
	v.Test = false
	v.HadWin = false
	v.HasInventory = false
	poolShowBenefitDto.Put(v)
}
