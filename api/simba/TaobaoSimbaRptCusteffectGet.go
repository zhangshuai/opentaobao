package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCusteffectGet 用户账户报表效果数据查询（只有汇总数据，无分类数据）
// taobao.simba.rpt.custeffect.get
//
// 用户账户报表效果数据查询（只有汇总数据，无分类数据）
func TaobaoSimbaRptCusteffectGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptCusteffectGetAPIRequest, resp *simba.TaobaoSimbaRptCusteffectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
