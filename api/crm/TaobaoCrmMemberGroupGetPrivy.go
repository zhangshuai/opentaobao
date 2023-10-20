package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmMemberGroupGetPrivy 获取买家身上的标签（隐私号版）
// taobao.crm.member.group.get.privy
//
// 获取买家身上的标签，不返回标签的总人数
func TaobaoCrmMemberGroupGetPrivy(clt *core.SDKClient, req *crm.TaobaoCrmMemberGroupGetPrivyAPIRequest, resp *crm.TaobaoCrmMemberGroupGetPrivyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
