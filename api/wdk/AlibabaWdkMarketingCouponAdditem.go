package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingCouponAdditem 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// alibaba.wdk.marketing.coupon.additem
//
// 在优惠券活动下添加商品。【注意，此接口暂不支持并发！】
// 如果是商品券，则添加的商品为券适用的商品；
// 如果是品类券，则添加的商品为券排除的商品；
func AlibabaWdkMarketingCouponAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponAdditemAPIRequest, resp *wdk.AlibabaWdkMarketingCouponAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
