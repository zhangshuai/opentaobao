package moziacl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moziacl"
)

// AlibabaMoziAclRoleAddPermissions 角色添加功能权限
// alibaba.mozi.acl.role.add.permissions
//
// 往角色中添加一批功能权限
func AlibabaMoziAclRoleAddPermissions(ctx context.Context, clt *core.SDKClient, req *moziacl.AlibabaMoziAclRoleAddPermissionsAPIRequest, resp *moziacl.AlibabaMoziAclRoleAddPermissionsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
