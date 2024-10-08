package degoperation

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// TaobaoDegoperationGetByEventkey 通用用户抽奖次数限制
// taobao.degoperation.get.by.eventkey
//
// 通用用户抽奖次数限制
func TaobaoDegoperationGetByEventkey(ctx context.Context, clt *core.SDKClient, req *degoperation.TaobaoDegoperationGetByEventkeyAPIRequest, resp *degoperation.TaobaoDegoperationGetByEventkeyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
