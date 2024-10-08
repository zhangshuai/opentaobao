package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScInvitecodeGet 淘宝客-公用-私域用户邀请码生成
// taobao.tbk.sc.invitecode.get
//
// 私域用户管理(即渠道管理或会员运营管理)功能中，通过此API可生成淘宝客自身的邀请码。
func TaobaoTbkScInvitecodeGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScInvitecodeGetAPIRequest, resp *tbk.TaobaoTbkScInvitecodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
