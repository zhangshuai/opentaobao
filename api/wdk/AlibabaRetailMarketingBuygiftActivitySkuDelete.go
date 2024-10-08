package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivitySkuDelete 删除单品买赠活动商品
// alibaba.retail.marketing.buygift.activity.sku.delete
//
// 删除单品买赠活动商品信息【同城零售】
func AlibabaRetailMarketingBuygiftActivitySkuDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
