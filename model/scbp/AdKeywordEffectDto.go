package scbp

import (
	"sync"
)

// AdKeywordEffectDto 结构体
type AdKeywordEffectDto struct {
	// 点击量
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 平均点击花费
	ClickCostAvg string `json:"click_cost_avg,omitempty" xml:"click_cost_avg,omitempty"`
	// 单位元，保留两位小数, 例如3.75表示3.75元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 百分比，保留两位小数，例如3.75表示3.75%
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 曝光
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 单位小时，保留一位小数，例如13.5表示13.5小时
	OnlineTime string `json:"online_time,omitempty" xml:"online_time,omitempty"`
}

var poolAdKeywordEffectDto = sync.Pool{
	New: func() any {
		return new(AdKeywordEffectDto)
	},
}

// GetAdKeywordEffectDto() 从对象池中获取AdKeywordEffectDto
func GetAdKeywordEffectDto() *AdKeywordEffectDto {
	return poolAdKeywordEffectDto.Get().(*AdKeywordEffectDto)
}

// ReleaseAdKeywordEffectDto 释放AdKeywordEffectDto
func ReleaseAdKeywordEffectDto(v *AdKeywordEffectDto) {
	v.ClickCnt = ""
	v.ClickCostAvg = ""
	v.Cost = ""
	v.Ctr = ""
	v.ImpressionCnt = ""
	v.Keyword = ""
	v.OnlineTime = ""
	poolAdKeywordEffectDto.Put(v)
}
