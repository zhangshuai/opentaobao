package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMessageSubscriptionQuery 查询用户是否有模版ID权限
// alitrip.merchant.galaxy.message.subscription.query
//
// 只是查询用户是否拥有权限ID
func AlitripMerchantGalaxyMessageSubscriptionQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
