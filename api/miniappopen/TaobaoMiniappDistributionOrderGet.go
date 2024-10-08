package miniappopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniappopen"
)

// TaobaoMiniappDistributionOrderGet 小程序投放-查询小程序投放计划信息
// taobao.miniapp.distribution.order.get
//
// 服务商可通过该API，读取自己开发的小程序对应的投放计划的相关信息
func TaobaoMiniappDistributionOrderGet(ctx context.Context, clt *core.SDKClient, req *miniappopen.TaobaoMiniappDistributionOrderGetAPIRequest, resp *miniappopen.TaobaoMiniappDistributionOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
