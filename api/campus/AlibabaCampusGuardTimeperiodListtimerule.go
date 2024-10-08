package campus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardTimeperiodListtimerule 门禁控制器查询时间规则
// alibaba.campus.guard.timeperiod.listtimerule
//
// 门禁控制器查询时间规则
func AlibabaCampusGuardTimeperiodListtimerule(ctx context.Context, clt *core.SDKClient, req *campus.AlibabaCampusGuardTimeperiodListtimeruleAPIRequest, resp *campus.AlibabaCampusGuardTimeperiodListtimeruleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
