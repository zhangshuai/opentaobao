package examination

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// AlibabaAlihealthMedicalOrderRefund 退款接口
// alibaba.alihealth.medical.order.refund
//
// 退款接口
func AlibabaAlihealthMedicalOrderRefund(ctx context.Context, clt *core.SDKClient, req *examination.AlibabaAlihealthMedicalOrderRefundAPIRequest, resp *examination.AlibabaAlihealthMedicalOrderRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
