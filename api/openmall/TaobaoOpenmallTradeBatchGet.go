package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeBatchGet 批量获取openmall订单
// taobao.openmall.trade.batch.get
//
// 批量获取openmall订单
// 注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
func TaobaoOpenmallTradeBatchGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeBatchGetAPIRequest, resp *openmall.TaobaoOpenmallTradeBatchGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
