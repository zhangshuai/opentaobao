package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiUserQuickBind 精灵用户绑定第三方账号信息
// alibaba.ai.user.quick.bind
//
// 人工智能实验室精灵用户绑定第三方账号信息接口，开放给Iot厂商做为厂商上送第三方账号信息的接口
func AlibabaAiUserQuickBind(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAiUserQuickBindAPIRequest, resp *tmallgenie.AlibabaAiUserQuickBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
