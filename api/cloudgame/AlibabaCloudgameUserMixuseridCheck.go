package cloudgame

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameUserMixuseridCheck 云游戏混淆用户ID校验
// alibaba.cloudgame.user.mixuserid.check
//
// 验证混淆用户ID是否合法
func AlibabaCloudgameUserMixuseridCheck(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameUserMixuseridCheckAPIRequest, resp *cloudgame.AlibabaCloudgameUserMixuseridCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
