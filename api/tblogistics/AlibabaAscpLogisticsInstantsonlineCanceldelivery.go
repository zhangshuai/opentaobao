package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlineCanceldelivery 同城配送在线下单取消下单取消呼叫的运力
// alibaba.ascp.logistics.instantsonline.canceldelivery
//
// 同城配送在线下单取消下单取消呼叫的运力
func AlibabaAscpLogisticsInstantsonlineCanceldelivery(ctx context.Context, clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIRequest, resp *tblogistics.AlibabaAscpLogisticsInstantsonlineCanceldeliveryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
