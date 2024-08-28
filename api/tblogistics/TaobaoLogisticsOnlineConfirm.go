package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsOnlineConfirm 确认发货通知接口
// taobao.logistics.online.confirm
//
// &lt;br&gt;&lt;font color=&#39;red&#39;&gt;仅在使用taobao.logistics.online.send 发货时未输入运单号的情况下，需要使用该接口补充填写运单号，来确认发货。&lt;br&gt;
// 确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】。&lt;/font&gt;
func TaobaoLogisticsOnlineConfirm(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsOnlineConfirmAPIRequest, resp *tblogistics.TaobaoLogisticsOnlineConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
