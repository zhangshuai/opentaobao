package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyUserRisk 查询微信账号的风险等级
// alitrip.merchant.galaxy.user.risk
//
// 【星河服务】查询微信账号的风险等级
func AlitripMerchantGalaxyUserRisk(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyUserRiskAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyUserRiskAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
