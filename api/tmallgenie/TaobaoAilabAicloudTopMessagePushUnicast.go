package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoAilabAicloudTopMessagePushUnicast 天猫精灵消息中心单播推送消息接口
// taobao.ailab.aicloud.top.message.push.unicast
//
// 天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。
func TaobaoAilabAicloudTopMessagePushUnicast(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoAilabAicloudTopMessagePushUnicastAPIRequest, resp *tmallgenie.TaobaoAilabAicloudTopMessagePushUnicastAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
