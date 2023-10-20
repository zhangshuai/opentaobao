package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// Taobaosubuserfullinfoget 获取指定账户子账号的详细信息
// taobao.subuser.fullinfo.get
//
// 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
func Taobaosubuserfullinfoget(clt *core.SDKClient, req *subuser.TaobaosubuserfullinfogetAPIRequest, session string) (*subuser.TaobaosubuserfullinfogetAPIResponse, error) {
	var resp subuser.TaobaosubuserfullinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
