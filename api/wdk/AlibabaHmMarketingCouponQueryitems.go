package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingCouponQueryitems 查询优惠券活动下的商品
// alibaba.hm.marketing.coupon.queryitems
//
// 查询优惠券活动下面的商品
func AlibabaHmMarketingCouponQueryitems(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingCouponQueryitemsAPIRequest, resp *wdk.AlibabaHmMarketingCouponQueryitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
