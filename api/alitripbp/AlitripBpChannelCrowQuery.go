package alitripbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripbp"
)

// AlitripBpChannelCrowQuery 人群匹配
// alitrip.bp.channel.crow.query
//
// 判断用户是否在圈选的人群中
func AlitripBpChannelCrowQuery(ctx context.Context, clt *core.SDKClient, req *alitripbp.AlitripBpChannelCrowQueryAPIRequest, resp *alitripbp.AlitripBpChannelCrowQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
