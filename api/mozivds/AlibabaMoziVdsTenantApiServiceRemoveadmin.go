package mozivds

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozivds"
)

// AlibabaMoziVdsTenantApiServiceRemoveadmin 删除租户管理员服务
// alibaba.mozi.vds.tenant.api.service.removeadmin
//
// 删除租户管理员top服务
func AlibabaMoziVdsTenantApiServiceRemoveadmin(clt *core.SDKClient, req *mozivds.AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest, resp *mozivds.AlibabaMoziVdsTenantApiServiceRemoveadminAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
