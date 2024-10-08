package car

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/car"
)

// TaobaoAlitripCarRentOrderCancel 租车-取消订单
// taobao.alitrip.car.rent.order.cancel
//
// 服务商主动取消用户订单或者拒绝取消订单.
func TaobaoAlitripCarRentOrderCancel(ctx context.Context, clt *core.SDKClient, req *car.TaobaoAlitripCarRentOrderCancelAPIRequest, resp *car.TaobaoAlitripCarRentOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
