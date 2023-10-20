package alihouse

import (
	"sync"
)

// BaseLoopLineDto 结构体
type BaseLoopLineDto struct {
	// 环线名称
	LoopLineName string `json:"loop_line_name,omitempty" xml:"loop_line_name,omitempty"`
	// 环线电子围栏
	LoopLineFence string `json:"loop_line_fence,omitempty" xml:"loop_line_fence,omitempty"`
	// 外部ID - 唯一
	OuterLineId string `json:"outer_line_id,omitempty" xml:"outer_line_id,omitempty"`
	// 是否删除 1 是 0 否
	IsDeleted string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 序号
	Number int64 `json:"number,omitempty" xml:"number,omitempty"`
	// 环线ID
	LoopLineId int64 `json:"loop_line_id,omitempty" xml:"loop_line_id,omitempty"`
	// 数据源类型（1-新房 2-二手房）
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
}

var poolBaseLoopLineDto = sync.Pool{
	New: func() any {
		return new(BaseLoopLineDto)
	},
}

// GetBaseLoopLineDto() 从对象池中获取BaseLoopLineDto
func GetBaseLoopLineDto() *BaseLoopLineDto {
	return poolBaseLoopLineDto.Get().(*BaseLoopLineDto)
}

// ReleaseBaseLoopLineDto 释放BaseLoopLineDto
func ReleaseBaseLoopLineDto(v *BaseLoopLineDto) {
	v.LoopLineName = ""
	v.LoopLineFence = ""
	v.OuterLineId = ""
	v.IsDeleted = ""
	v.CityId = 0
	v.Number = 0
	v.LoopLineId = 0
	v.SourceType = 0
	poolBaseLoopLineDto.Put(v)
}
