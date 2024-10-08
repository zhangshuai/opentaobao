package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillBoxPostBackBox RT收箱回传
// alibaba.wdk.fulfill.box.post.back.box
//
// RT收箱后，信息同步履约，履约同通知UMS 容器管理
func AlibabaWdkFulfillBoxPostBackBox(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillBoxPostBackBoxAPIRequest, resp *wdk.AlibabaWdkFulfillBoxPostBackBoxAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
