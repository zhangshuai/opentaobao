package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivityCreate 创建单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.create
//
// 同城零售单品特价活动创建
func AlibabaRetailMarketingItemdiscountActivityCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivityCreateAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivityCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
