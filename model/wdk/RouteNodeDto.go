package wdk

import (
	"sync"
)

// RouteNodeDto 结构体
type RouteNodeDto struct {
	// 节点类型
	NodeType string `json:"node_type,omitempty" xml:"node_type,omitempty"`
	// 节点名称
	NodeCode string `json:"node_code,omitempty" xml:"node_code,omitempty"`
	// 节点序号
	NodeIndex int64 `json:"node_index,omitempty" xml:"node_index,omitempty"`
}

var poolRouteNodeDto = sync.Pool{
	New: func() any {
		return new(RouteNodeDto)
	},
}

// GetRouteNodeDto() 从对象池中获取RouteNodeDto
func GetRouteNodeDto() *RouteNodeDto {
	return poolRouteNodeDto.Get().(*RouteNodeDto)
}

// ReleaseRouteNodeDto 释放RouteNodeDto
func ReleaseRouteNodeDto(v *RouteNodeDto) {
	v.NodeType = ""
	v.NodeCode = ""
	v.NodeIndex = 0
	poolRouteNodeDto.Put(v)
}
