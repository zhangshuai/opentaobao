package tmc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmc"
)

// TaobaoTmcUserGet 获取用户已开通消息
// taobao.tmc.user.get
//
// 查询指定用户开通的消息通道和组
func TaobaoTmcUserGet(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcUserGetAPIRequest, resp *tmc.TaobaoTmcUserGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
