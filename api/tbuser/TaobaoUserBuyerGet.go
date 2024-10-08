package tbuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserBuyerGet 查询买家信息API
// taobao.user.buyer.get
//
// 查询买家信息API，只能买家类应用调用。
func TaobaoUserBuyerGet(ctx context.Context, clt *core.SDKClient, req *tbuser.TaobaoUserBuyerGetAPIRequest, resp *tbuser.TaobaoUserBuyerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
