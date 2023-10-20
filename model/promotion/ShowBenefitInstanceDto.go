package promotion

import (
	"sync"
)

// ShowBenefitInstanceDto 结构体
type ShowBenefitInstanceDto struct {
	// 权益code
	BenefitCode string `json:"benefit_code,omitempty" xml:"benefit_code,omitempty"`
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 中奖记录ID
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
}

var poolShowBenefitInstanceDto = sync.Pool{
	New: func() any {
		return new(ShowBenefitInstanceDto)
	},
}

// GetShowBenefitInstanceDto() 从对象池中获取ShowBenefitInstanceDto
func GetShowBenefitInstanceDto() *ShowBenefitInstanceDto {
	return poolShowBenefitInstanceDto.Get().(*ShowBenefitInstanceDto)
}

// ReleaseShowBenefitInstanceDto 释放ShowBenefitInstanceDto
func ReleaseShowBenefitInstanceDto(v *ShowBenefitInstanceDto) {
	v.BenefitCode = ""
	v.BenefitType = ""
	v.RecordId = 0
	poolShowBenefitInstanceDto.Put(v)
}
