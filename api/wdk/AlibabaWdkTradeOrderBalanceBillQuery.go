package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTradeOrderBalanceBillQuery 分页拉取订单数据
// alibaba.wdk.trade.order.balance.bill.query
//
// 提供接口供外部调用，分页拉取订单数据
func AlibabaWdkTradeOrderBalanceBillQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkTradeOrderBalanceBillQueryAPIRequest, resp *wdk.AlibabaWdkTradeOrderBalanceBillQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
