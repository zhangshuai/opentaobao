package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpensecurityIsvUidGet 获取open security uid for isv
// taobao.opensecurity.isv.uid.get
//
// 根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
func TaobaoOpensecurityIsvUidGet(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpensecurityIsvUidGetAPIRequest, resp *user.TaobaoOpensecurityIsvUidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
