package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceFlowRenew 工单(入驻、加盘、续约)续约
// alibaba.einvoice.flow.renew
//
// 工单(含入驻、加盘、续约工单)续约能力开放
func AlibabaEinvoiceFlowRenew(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceFlowRenewAPIRequest, resp *einvoice.AlibabaEinvoiceFlowRenewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
