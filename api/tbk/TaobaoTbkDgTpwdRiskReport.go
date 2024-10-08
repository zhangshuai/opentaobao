package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgTpwdRiskReport 淘宝客-推广者-淘口令预警及拦截查询
// taobao.tbk.dg.tpwd.risk.report
//
// 淘宝客-推广者-淘口令预警及拦截查询
func TaobaoTbkDgTpwdRiskReport(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgTpwdRiskReportAPIRequest, resp *tbk.TaobaoTbkDgTpwdRiskReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
