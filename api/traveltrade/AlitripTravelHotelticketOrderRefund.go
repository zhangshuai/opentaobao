package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelHotelticketOrderRefund 退款结结果通知
// alitrip.travel.hotelticket.order.refund
//
// 退款结果通知
func AlitripTravelHotelticketOrderRefund(clt *core.SDKClient, req *traveltrade.AlitripTravelHotelticketOrderRefundAPIRequest, resp *traveltrade.AlitripTravelHotelticketOrderRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
