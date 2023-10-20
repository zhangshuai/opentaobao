package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainHotelSearch 【商旅】酒店订单查询
// alitrip.btrip.supplychain.hotel.search
//
// 【商旅】酒店订单查询
func AlitripBtripSupplychainHotelSearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainHotelSearchAPIRequest, resp *btrip.AlitripBtripSupplychainHotelSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
