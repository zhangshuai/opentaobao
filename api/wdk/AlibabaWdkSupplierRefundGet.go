package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSupplierRefundGet 五道口按订单号批量查询供应商退款单
// alibaba.wdk.supplier.refund.get
//
// 五道口按订单号批量查询供应商退款单
func AlibabaWdkSupplierRefundGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSupplierRefundGetAPIRequest, resp *wdk.AlibabaWdkSupplierRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
