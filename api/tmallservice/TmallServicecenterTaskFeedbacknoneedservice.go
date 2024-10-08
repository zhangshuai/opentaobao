package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterTaskFeedbacknoneedservice 服务商反馈无需安装工单接口
// tmall.servicecenter.task.feedbacknoneedservice
//
// 服务商反馈无需安装工单接口
func TmallServicecenterTaskFeedbacknoneedservice(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterTaskFeedbacknoneedserviceAPIRequest, resp *tmallservice.TmallServicecenterTaskFeedbacknoneedserviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
