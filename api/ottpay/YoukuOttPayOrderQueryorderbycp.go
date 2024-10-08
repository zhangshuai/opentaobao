package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderQueryorderbycp 订单查询接口(cp订单号查询)
// youku.ott.pay.order.queryorderbycp
//
// 给商户服务端查询订单状态
func YoukuOttPayOrderQueryorderbycp(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttPayOrderQueryorderbycpAPIRequest, resp *ottpay.YoukuOttPayOrderQueryorderbycpAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
