package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolActivitySkuDelete 删除商品池活动商品【同城零售】
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
func AlibabaRetailMarketingItempoolActivitySkuDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
