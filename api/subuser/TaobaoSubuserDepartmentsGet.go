package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubuserDepartmentsGet 获取指定账户的所有部门列表
// taobao.subuser.departments.get
//
// 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。
func TaobaoSubuserDepartmentsGet(clt *core.SDKClient, req *subuser.TaobaoSubuserDepartmentsGetAPIRequest, resp *subuser.TaobaoSubuserDepartmentsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
