package icbu

import (
	"sync"
)

// ProductGroup 结构体
type ProductGroup struct {
	// 下级分组ID列表
	ChildrenIdList []int64 `json:"children_id_list,omitempty" xml:"children_id_list>int64,omitempty"`
	// 分组名称
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 上级分组ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 分组ID
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// 父节点id，父节点处在分组树的二级
	ParentId2 int64 `json:"parent_id2,omitempty" xml:"parent_id2,omitempty"`
}

var poolProductGroup = sync.Pool{
	New: func() any {
		return new(ProductGroup)
	},
}

// GetProductGroup() 从对象池中获取ProductGroup
func GetProductGroup() *ProductGroup {
	return poolProductGroup.Get().(*ProductGroup)
}

// ReleaseProductGroup 释放ProductGroup
func ReleaseProductGroup(v *ProductGroup) {
	v.ChildrenIdList = v.ChildrenIdList[:0]
	v.GroupName = ""
	v.ParentId = 0
	v.GroupId = 0
	v.ParentId2 = 0
	poolProductGroup.Put(v)
}
