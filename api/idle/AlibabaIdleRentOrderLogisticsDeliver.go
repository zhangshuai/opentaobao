package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderLogisticsDeliver 创建揽收物流
// alibaba.idle.rent.order.logistics.deliver
//
// 创建揽收物流
// 商家去物流公司创建物流订单
func AlibabaIdleRentOrderLogisticsDeliver(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentOrderLogisticsDeliverAPIRequest, resp *idle.AlibabaIdleRentOrderLogisticsDeliverAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
