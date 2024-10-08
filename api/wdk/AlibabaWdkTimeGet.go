package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkTimeGet 获得当前系统时间
// alibaba.wdk.time.get
//
// 获得当前系统时间
func AlibabaWdkTimeGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkTimeGetAPIRequest, resp *wdk.AlibabaWdkTimeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
