package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoLogisticsExpressCollectSync 服饰逆向揽收信息同步
// taobao.logistics.express.collect.sync
//
// 服饰逆向揽收信息同步
func TaobaoLogisticsExpressCollectSync(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoLogisticsExpressCollectSyncAPIRequest, resp *logistic.TaobaoLogisticsExpressCollectSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
