package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingDiscountItemAddAsync 特价批量新增商品
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
func AlibabaHmMarketingDiscountItemAddAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingDiscountItemAddAsyncAPIRequest, resp *wdk.AlibabaHmMarketingDiscountItemAddAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
