package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectTradeOrder 旺铺楼盘和交易商品排序
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
func AlibabaAlihouseNewhomeProjectTradeOrder(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
