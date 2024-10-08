package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// TaobaoRdcAligeniusWarehouseReverseEventUpdate 销退单事件回传接口
// taobao.rdc.aligenius.warehouse.reverse.event.update
//
// 用于erp回传销退单相关信息到平台
func TaobaoRdcAligeniusWarehouseReverseEventUpdate(ctx context.Context, clt *core.SDKClient, req *logistic.TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIRequest, resp *logistic.TaobaoRdcAligeniusWarehouseReverseEventUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
