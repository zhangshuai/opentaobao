package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingExpirePromotionDelete 短保优惠删除
// alibaba.wdk.marketing.expire.promotion.delete
//
// 短保优惠删除
func AlibabaWdkMarketingExpirePromotionDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionDeleteAPIRequest, resp *wdk.AlibabaWdkMarketingExpirePromotionDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
