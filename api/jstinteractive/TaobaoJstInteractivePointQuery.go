package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractivePointQuery 互动积分查询接口
// taobao.jst.interactive.point.query
//
// 查询用户的互动积分
func TaobaoJstInteractivePointQuery(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointQueryAPIRequest, resp *jstinteractive.TaobaoJstInteractivePointQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
