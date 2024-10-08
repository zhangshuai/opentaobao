package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceSerialnoBatchGenerate 开票流水号批量生成接口
// alibaba.einvoice.serialno.batch.generate
//
// 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
// 优先使用alibaba.einvoice.serial.generate。
func AlibabaEinvoiceSerialnoBatchGenerate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceSerialnoBatchGenerateAPIRequest, resp *einvoice.AlibabaEinvoiceSerialnoBatchGenerateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
