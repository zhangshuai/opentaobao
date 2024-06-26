package tbk

import (
	"sync"
)

// SubsidyDetailDto 结构体
type SubsidyDetailDto struct {
	// 该笔订单包含的补贴类型
	SubsidyType string `json:"subsidy_type,omitempty" xml:"subsidy_type,omitempty"`
	// 补贴比率
	SubsidyRate float64 `json:"subsidy_rate,omitempty" xml:"subsidy_rate,omitempty"`
	// 对应补贴类型的补贴金额
	SubsidyFee float64 `json:"subsidy_fee,omitempty" xml:"subsidy_fee,omitempty"`
	// 单笔订单补贴上限。对应补贴类型的单笔订单补贴上限
	SubsidyUpperLimit float64 `json:"subsidy_upper_limit,omitempty" xml:"subsidy_upper_limit,omitempty"`
	// 补贴分成比率
	SubsidyShareRate float64 `json:"subsidy_share_rate,omitempty" xml:"subsidy_share_rate,omitempty"`
}

var poolSubsidyDetailDto = sync.Pool{
	New: func() any {
		return new(SubsidyDetailDto)
	},
}

// GetSubsidyDetailDto() 从对象池中获取SubsidyDetailDto
func GetSubsidyDetailDto() *SubsidyDetailDto {
	return poolSubsidyDetailDto.Get().(*SubsidyDetailDto)
}

// ReleaseSubsidyDetailDto 释放SubsidyDetailDto
func ReleaseSubsidyDetailDto(v *SubsidyDetailDto) {
	v.SubsidyType = ""
	v.SubsidyRate = 0
	v.SubsidyFee = 0
	v.SubsidyUpperLimit = 0
	v.SubsidyShareRate = 0
	poolSubsidyDetailDto.Put(v)
}
