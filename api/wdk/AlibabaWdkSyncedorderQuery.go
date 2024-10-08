package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSyncedorderQuery 五道口查询同步订单
// alibaba.wdk.syncedorder.query
//
// 外部商户查询同步到五道口的订单
func AlibabaWdkSyncedorderQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSyncedorderQueryAPIRequest, resp *wdk.AlibabaWdkSyncedorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
