package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoMobilePromotionCouponSellerSearch 查询绑定卖家优惠券相关信息(手淘专用)
// taobao.mobile.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方
func TaobaoMobilePromotionCouponSellerSearch(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoMobilePromotionCouponSellerSearchAPIRequest, resp *promotion.TaobaoMobilePromotionCouponSellerSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
