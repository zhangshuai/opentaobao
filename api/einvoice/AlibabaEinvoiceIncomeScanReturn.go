package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeScanReturn 进项扫描状态回传
// alibaba.einvoice.income.scan.return
//
// 回传进项扫描每个阶段的状态，比如ocr开始，ocr结束，查验开始，查验结束等
func AlibabaEinvoiceIncomeScanReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeScanReturnAPIRequest, resp *einvoice.AlibabaEinvoiceIncomeScanReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
