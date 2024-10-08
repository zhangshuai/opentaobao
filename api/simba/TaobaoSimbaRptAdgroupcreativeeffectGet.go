package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupcreativeeffectGet 推广组下的创意报表效果数据查询(汇总数据，不分类型)
// taobao.simba.rpt.adgroupcreativeeffect.get
//
// 推广组下的创意报表效果数据查询(汇总数据，不分类型)
func TaobaoSimbaRptAdgroupcreativeeffectGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
