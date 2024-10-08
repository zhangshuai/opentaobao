package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderWorkCallback 仓配作业结果回传接口
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
func AlibabaAelophyOrderWorkCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyOrderWorkCallbackAPIRequest, resp *wdk.AlibabaAelophyOrderWorkCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
