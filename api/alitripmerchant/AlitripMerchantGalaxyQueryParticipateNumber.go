package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyQueryParticipateNumber 星河-抽奖活动次数查询
// alitrip.merchant.galaxy.query.participate.number
//
// 雅高小程序抽奖活动，查询用户抽奖次数
func AlitripMerchantGalaxyQueryParticipateNumber(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyQueryParticipateNumberAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyQueryParticipateNumberAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
