package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkUmsOutboundSortingCallbackforpulltask dps分货-任务拉取确定接口
// wdk.ums.outbound.sorting.callbackforpulltask
//
// dps分货-任务拉取确定接口
func WdkUmsOutboundSortingCallbackforpulltask(ctx context.Context, clt *core.SDKClient, req *wdk.WdkUmsOutboundSortingCallbackforpulltaskAPIRequest, resp *wdk.WdkUmsOutboundSortingCallbackforpulltaskAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
