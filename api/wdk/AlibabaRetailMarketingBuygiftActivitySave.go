package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivitySave 【同城零售】单品买赠活动保存
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
func AlibabaRetailMarketingBuygiftActivitySave(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivitySaveAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivitySaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
