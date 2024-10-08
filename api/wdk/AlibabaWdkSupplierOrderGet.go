package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSupplierOrderGet 五道口按订单号批量查询供应商正向订单
// alibaba.wdk.supplier.order.get
//
// 五道口按订单号批量查询供应商正向订单
func AlibabaWdkSupplierOrderGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSupplierOrderGetAPIRequest, resp *wdk.AlibabaWdkSupplierOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
