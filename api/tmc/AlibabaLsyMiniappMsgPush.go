package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// AlibabaLsyMiniappMsgPush 零售云小程序消息推送
// alibaba.lsy.miniapp.msg.push
//
// 零售云小程序消息推送，推送消息至零售云（喵零等）
func AlibabaLsyMiniappMsgPush(ctx context.Context, clt *core.SDKClient, req *tmc.AlibabaLsyMiniappMsgPushAPIRequest, resp *tmc.AlibabaLsyMiniappMsgPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
