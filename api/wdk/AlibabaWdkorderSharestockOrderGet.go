package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkorderSharestockOrderGet 猫超商户订单拉取
// alibaba.wdkorder.sharestock.order.get
//
// 商户拉取猫超订单数据
func AlibabaWdkorderSharestockOrderGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkorderSharestockOrderGetAPIRequest, resp *wdk.AlibabaWdkorderSharestockOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
