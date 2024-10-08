package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardExtrachargeCreate 创建工单额外收费项
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
func TmallServicecenterWorkcardExtrachargeCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardExtrachargeCreateAPIRequest, resp *tmallservice.TmallServicecenterWorkcardExtrachargeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
