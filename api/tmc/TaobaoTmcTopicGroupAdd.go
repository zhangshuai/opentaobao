package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcTopicGroupAdd topic分组路由
// taobao.tmc.topic.group.add
//
// 根据topic名称路由消息到不同的分组。（前提：发送方未指定分组名）
// 如果是需要授权的消息，分组路由先判断用户分组路由(使用taobao.tmc.group.add添加的路由)，用户分组路由不存在时，才会判断topic分组路由
func TaobaoTmcTopicGroupAdd(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcTopicGroupAddAPIRequest, resp *tmc.TaobaoTmcTopicGroupAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
