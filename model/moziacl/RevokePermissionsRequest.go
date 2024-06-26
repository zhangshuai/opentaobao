package moziacl

import (
	"sync"
)

// RevokePermissionsRequest 结构体
type RevokePermissionsRequest struct {
	// 回收权限的name列表
	PermissionNames []string `json:"permission_names,omitempty" xml:"permission_names>string,omitempty"`
	// 回收权限所在的app name
	TargetAppName string `json:"target_app_name,omitempty" xml:"target_app_name,omitempty"`
	// 请求扩展字段
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 回收主体对象
	Principal *BucUserPrincipalParam `json:"principal,omitempty" xml:"principal,omitempty"`
}

var poolRevokePermissionsRequest = sync.Pool{
	New: func() any {
		return new(RevokePermissionsRequest)
	},
}

// GetRevokePermissionsRequest() 从对象池中获取RevokePermissionsRequest
func GetRevokePermissionsRequest() *RevokePermissionsRequest {
	return poolRevokePermissionsRequest.Get().(*RevokePermissionsRequest)
}

// ReleaseRevokePermissionsRequest 释放RevokePermissionsRequest
func ReleaseRevokePermissionsRequest(v *RevokePermissionsRequest) {
	v.PermissionNames = v.PermissionNames[:0]
	v.TargetAppName = ""
	v.RequestMetaData = ""
	v.Principal = nil
	poolRevokePermissionsRequest.Put(v)
}
