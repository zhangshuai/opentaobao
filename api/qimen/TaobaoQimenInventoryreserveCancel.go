package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryreserveCancel 库存预占取消接口
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
func TaobaoQimenInventoryreserveCancel(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenInventoryreserveCancelAPIRequest, resp *qimen.TaobaoQimenInventoryreserveCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
