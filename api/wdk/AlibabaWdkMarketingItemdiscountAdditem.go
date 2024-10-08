package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountAdditem 报名特价商品
// alibaba.wdk.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
func AlibabaWdkMarketingItemdiscountAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountAdditemAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
