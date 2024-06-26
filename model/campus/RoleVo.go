package campus

import (
	"sync"
)

// RoleVo 结构体
type RoleVo struct {
	// 角色id
	RoleId string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// 角色名称
	RoleName string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// 角色类型(admin代表人员类型,device代表设备类型的角色)
	RoleType string `json:"role_type,omitempty" xml:"role_type,omitempty"`
}

var poolRoleVo = sync.Pool{
	New: func() any {
		return new(RoleVo)
	},
}

// GetRoleVo() 从对象池中获取RoleVo
func GetRoleVo() *RoleVo {
	return poolRoleVo.Get().(*RoleVo)
}

// ReleaseRoleVo 释放RoleVo
func ReleaseRoleVo(v *RoleVo) {
	v.RoleId = ""
	v.RoleName = ""
	v.RoleType = ""
	poolRoleVo.Put(v)
}
