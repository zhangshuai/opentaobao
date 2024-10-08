package alscmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alscmerchant"
)

// AlibabaAlscMerchantExtTicketcodeUse 外部核销服务
// alibaba.alsc.merchant.ext.ticketcode.use
//
// 外部核销服务
func AlibabaAlscMerchantExtTicketcodeUse(ctx context.Context, clt *core.SDKClient, req *alscmerchant.AlibabaAlscMerchantExtTicketcodeUseAPIRequest, resp *alscmerchant.AlibabaAlscMerchantExtTicketcodeUseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
