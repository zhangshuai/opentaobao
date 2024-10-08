package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoKoubeiSaasBaseOperationConfigSync 商家基础经营设置信息同步
// taobao.koubei.saas.base.operation.config.sync
//
// ISV接入口碑SAAS后, 经营设置数据同步到口碑SAAS
func TaobaoKoubeiSaasBaseOperationConfigSync(ctx context.Context, clt *core.SDKClient, req *alsc.TaobaoKoubeiSaasBaseOperationConfigSyncAPIRequest, resp *alsc.TaobaoKoubeiSaasBaseOperationConfigSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
