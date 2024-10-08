package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeaseCitysynchronize 天猫开新车租后分期城市信息同步
// tmall.car.lease.citysynchronize
//
// 天猫开新车租后分期城市信息同步
func TmallCarLeaseCitysynchronize(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeaseCitysynchronizeAPIRequest, resp *tmallcar.TmallCarLeaseCitysynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
