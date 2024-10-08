package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoShopUpdate 更新店铺基本信息
// taobao.shop.update
//
// 目前只支持标题、公告和描述的更新
func TaobaoShopUpdate(ctx context.Context, clt *core.SDKClient, req *shop.TaobaoShopUpdateAPIRequest, resp *shop.TaobaoShopUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
