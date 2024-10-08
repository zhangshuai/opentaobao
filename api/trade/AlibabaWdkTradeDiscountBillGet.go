package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeDiscountBillGet 订单优惠账单查询
// alibaba.wdk.trade.discount.bill.get
//
// 商家查询订单优惠账单
func AlibabaWdkTradeDiscountBillGet(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaWdkTradeDiscountBillGetAPIRequest, resp *trade.AlibabaWdkTradeDiscountBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
