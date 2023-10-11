package promotion

// BenefitSelectorVo 结构体
type BenefitSelectorVo struct {
	// 详情list
	PackDetailList []BenefitTemplateVo `json:"pack_detail_list,omitempty" xml:"pack_detail_list>benefit_template_vo,omitempty"`
	// 权益名称
	BenefitName string `json:"benefit_name,omitempty" xml:"benefit_name,omitempty"`
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 权益类型id
	BenefitTypeLong int64 `json:"benefit_type_long,omitempty" xml:"benefit_type_long,omitempty"`
	// 选择器步骤选择2时，查询指定类型权益列表分页查询返回值，数据总量
	TemplateTotalCount int64 `json:"template_total_count,omitempty" xml:"template_total_count,omitempty"`
}
