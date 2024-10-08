package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaBaseOrderSupplierNotify 阿里通信运营商信息回传
// alibaba.base.order.supplier.notify
//
// 接收阿里通信流量运营商信息回传
func AlibabaBaseOrderSupplierNotify(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaBaseOrderSupplierNotifyAPIRequest, resp *alicom.AlibabaBaseOrderSupplierNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
