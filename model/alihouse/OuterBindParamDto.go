package alihouse

import (
	"sync"
)

// OuterBindParamDto 结构体
type OuterBindParamDto struct {
	// 外部工具唯一id，如购房登记ID
	OuterToolId string `json:"outer_tool_id,omitempty" xml:"outer_tool_id,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部货id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
}

var poolOuterBindParamDto = sync.Pool{
	New: func() any {
		return new(OuterBindParamDto)
	},
}

// GetOuterBindParamDto() 从对象池中获取OuterBindParamDto
func GetOuterBindParamDto() *OuterBindParamDto {
	return poolOuterBindParamDto.Get().(*OuterBindParamDto)
}

// ReleaseOuterBindParamDto 释放OuterBindParamDto
func ReleaseOuterBindParamDto(v *OuterBindParamDto) {
	v.OuterToolId = ""
	v.OuterId = ""
	v.OuterTid = ""
	v.OuterStoreId = ""
	poolOuterBindParamDto.Put(v)
}
