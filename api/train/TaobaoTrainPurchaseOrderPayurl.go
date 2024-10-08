package train

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// TaobaoTrainPurchaseOrderPayurl 火车票采购商接口-获取支付链接
// taobao.train.purchase.order.payurl
//
// 火车票采购商接口-获取支付链接
func TaobaoTrainPurchaseOrderPayurl(ctx context.Context, clt *core.SDKClient, req *train.TaobaoTrainPurchaseOrderPayurlAPIRequest, resp *train.TaobaoTrainPurchaseOrderPayurlAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
