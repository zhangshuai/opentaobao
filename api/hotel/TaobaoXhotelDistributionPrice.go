package hotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// TaobaoXhotelDistributionPrice 飞猪分销通用酒店报价接口
// taobao.xhotel.distribution.price
//
// 飞猪分销通用酒店报价接口
func TaobaoXhotelDistributionPrice(ctx context.Context, clt *core.SDKClient, req *hotel.TaobaoXhotelDistributionPriceAPIRequest, resp *hotel.TaobaoXhotelDistributionPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
