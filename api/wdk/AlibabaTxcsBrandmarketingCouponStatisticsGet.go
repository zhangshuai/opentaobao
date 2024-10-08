package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTxcsBrandmarketingCouponStatisticsGet 品牌营销导购员券推广统计数据回流
// alibaba.txcs.brandmarketing.coupon.statistics.get
//
// 请求券统计数据回流
func AlibabaTxcsBrandmarketingCouponStatisticsGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest, resp *wdk.AlibabaTxcsBrandmarketingCouponStatisticsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
