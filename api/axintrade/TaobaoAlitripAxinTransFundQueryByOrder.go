package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransFundQueryByOrder 通过外部订单ID查询所有资金单
// taobao.alitrip.axin.trans.fund.query.by.order
//
// 阿信供销平台-通过外部订单ID查询所有资金单
func TaobaoAlitripAxinTransFundQueryByOrder(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundQueryByOrderAPIRequest, resp *axintrade.TaobaoAlitripAxinTransFundQueryByOrderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
