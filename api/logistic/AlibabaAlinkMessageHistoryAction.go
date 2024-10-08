package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaAlinkMessageHistoryAction 操作历史消息
// alibaba.alink.message.history.action
//
// 阿里智能操作历史消息
func AlibabaAlinkMessageHistoryAction(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaAlinkMessageHistoryActionAPIRequest, resp *logistic.AlibabaAlinkMessageHistoryActionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
