package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderAlipayfaceSettle 信用住订单结账接口
// taobao.xhotel.order.alipayface.settle
//
// 用于离店付订单在客人离店后，发起结账以及扣款等后续动作
func TaobaoXhotelOrderAlipayfaceSettle(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceSettleAPIRequest, resp *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceSettleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
