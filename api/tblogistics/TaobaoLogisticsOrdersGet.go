package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsOrdersGet 批量查询物流订单
// taobao.logistics.orders.get
//
// 批量查询物流订单。
func TaobaoLogisticsOrdersGet(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsOrdersGetAPIRequest, resp *tblogistics.TaobaoLogisticsOrdersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
