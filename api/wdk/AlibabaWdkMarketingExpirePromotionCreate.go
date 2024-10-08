package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingExpirePromotionCreate 短保优惠创建
// alibaba.wdk.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
func AlibabaWdkMarketingExpirePromotionCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingExpirePromotionCreateAPIRequest, resp *wdk.AlibabaWdkMarketingExpirePromotionCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
