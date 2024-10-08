package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceAmountCheck 开票量核对接口
// alibaba.einvoice.amount.check
//
// 跟开票服务商核对历史开票量，用来对账
func AlibabaEinvoiceAmountCheck(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceAmountCheckAPIRequest, resp *einvoice.AlibabaEinvoiceAmountCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
