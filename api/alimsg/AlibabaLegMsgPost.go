package alimsg

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimsg"
)

// AlibabaLegMsgPost 集团法务消息发送
// alibaba.leg.msg.post
//
// 消息发送能力
func AlibabaLegMsgPost(ctx context.Context, clt *core.SDKClient, req *alimsg.AlibabaLegMsgPostAPIRequest, resp *alimsg.AlibabaLegMsgPostAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
