package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopHotelExceedapplyGet 商旅酒店第三方超标审批单搜索接口
// alitrip.btrip.corpop.hotel.exceedapply.get
//
// 商旅酒店第三方超标审批单搜索接口
func AlitripBtripCorpopHotelExceedapplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopHotelExceedapplyGetAPIRequest, resp *btrip.AlitripBtripCorpopHotelExceedapplyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
