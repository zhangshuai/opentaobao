package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivityCreate 创建商品池活动【同城零售】
// alibaba.retail.marketing.itempool.activity.create
//
// 同城零售商品池活动创建
func AlibabaRetailMarketingItempoolActivityCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityCreateAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
