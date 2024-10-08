package lstvending

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// AlibabaLstVendingGoodsSave 自动售卖机商品回传
// alibaba.lst.vending.goods.save
//
// 零售通自动售卖机商品数据回流。
func AlibabaLstVendingGoodsSave(ctx context.Context, clt *core.SDKClient, req *lstvending.AlibabaLstVendingGoodsSaveAPIRequest, resp *lstvending.AlibabaLstVendingGoodsSaveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
