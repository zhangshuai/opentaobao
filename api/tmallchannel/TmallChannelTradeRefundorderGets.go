package tmallchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TmallChannelTradeRefundorderGets 供应商查询退款单
// tmall.channel.trade.refundorder.gets
//
// 供应商分页查询退款单
func TmallChannelTradeRefundorderGets(ctx context.Context, clt *core.SDKClient, req *tmallchannel.TmallChannelTradeRefundorderGetsAPIRequest, resp *tmallchannel.TmallChannelTradeRefundorderGetsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
