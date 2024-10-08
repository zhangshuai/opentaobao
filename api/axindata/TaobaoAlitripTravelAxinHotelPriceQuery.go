package axindata

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelPriceQuery 阿信酒店分销-实时报价查询
// taobao.alitrip.travel.axin.hotel.price.query
//
// 酒店实时报价查询
func TaobaoAlitripTravelAxinHotelPriceQuery(ctx context.Context, clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest, resp *axindata.TaobaoAlitripTravelAxinHotelPriceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
