package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaDataCouponGet 获取优惠券信息
// alibaba.data.coupon.get
//
// 获取优惠券信息，仅作客户端鉴权虚拟api使用
func AlibabaDataCouponGet(ctx context.Context, clt *core.SDKClient, req *shop.AlibabaDataCouponGetAPIRequest, resp *shop.AlibabaDataCouponGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
