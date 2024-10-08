package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoJushitaJdpUsersGet 获取开通的订单同步服务的用户
// taobao.jushita.jdp.users.get
//
// 获取开通的订单同步服务的用户，含有rds的路由关系
func TaobaoJushitaJdpUsersGet(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJushitaJdpUsersGetAPIRequest, resp *jst.TaobaoJushitaJdpUsersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
