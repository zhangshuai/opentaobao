package moziacl

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclRoleRemovePermissions 角色移除功能权限
// alibaba.mozi.acl.role.remove.permissions
//
// 从角色中移除一批功能权限
func AlibabaMoziAclRoleRemovePermissions(clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleRemovePermissionsAPIRequest, resp *moziacl.AlibabaMoziAclRoleRemovePermissionsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
