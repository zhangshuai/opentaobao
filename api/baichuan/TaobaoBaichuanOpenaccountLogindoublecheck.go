package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// TaobaoBaichuanOpenaccountLogindoublecheck 百川登录二次验证
// taobao.baichuan.openaccount.logindoublecheck
//
// 百川登录二次验证
func TaobaoBaichuanOpenaccountLogindoublecheck(ctx context.Context, clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIRequest, resp *baichuan.TaobaoBaichuanOpenaccountLogindoublecheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
