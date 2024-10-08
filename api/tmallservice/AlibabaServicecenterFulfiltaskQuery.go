package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterFulfiltaskQuery 核销单查询
// alibaba.servicecenter.fulfiltask.query
//
// 当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务
func AlibabaServicecenterFulfiltaskQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.AlibabaServicecenterFulfiltaskQueryAPIRequest, resp *tmallservice.AlibabaServicecenterFulfiltaskQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
