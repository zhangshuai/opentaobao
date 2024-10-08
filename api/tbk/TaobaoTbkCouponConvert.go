package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkCouponConvert 淘宝客-推广者-单品券高效转链
// taobao.tbk.coupon.convert
//
// 单品券高效转链API
func TaobaoTbkCouponConvert(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkCouponConvertAPIRequest, resp *tbk.TaobaoTbkCouponConvertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
