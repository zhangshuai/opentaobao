package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAuto v5.0付费会员卡开票抬头自匹配
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.auto
//
// v5.0付费会员卡开票抬头自匹配
func AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAuto(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsAutoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
