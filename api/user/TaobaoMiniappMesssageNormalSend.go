package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoMiniappMesssageNormalSend 轻店铺下行普通消息给用户
// taobao.miniapp.messsage.normal.send
//
// 小程序下行单个普通消息
func TaobaoMiniappMesssageNormalSend(ctx context.Context, clt *core.SDKClient, req *user.TaobaoMiniappMesssageNormalSendAPIRequest, resp *user.TaobaoMiniappMesssageNormalSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
