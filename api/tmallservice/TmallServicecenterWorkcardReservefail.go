package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReservefail 预约失败
// tmall.servicecenter.workcard.reservefail
//
// 服务商调用该接口回传工单预约失败
func TmallServicecenterWorkcardReservefail(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReservefailAPIRequest, resp *tmallservice.TmallServicecenterWorkcardReservefailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
