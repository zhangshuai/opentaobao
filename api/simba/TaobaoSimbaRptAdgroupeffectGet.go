package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupeffectGet 推广组效果报表数据对象
// taobao.simba.rpt.adgroupeffect.get
//
// 推广组效果报表数据对象
func TaobaoSimbaRptAdgroupeffectGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupeffectGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupeffectGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
