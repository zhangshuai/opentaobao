package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemCarturlGet 加购URL获取
// taobao.item.carturl.get
//
// 获取加购URL，支持添加商品到购物车
func TaobaoItemCarturlGet(ctx context.Context, clt *core.SDKClient, req *product.TaobaoItemCarturlGetAPIRequest, resp *product.TaobaoItemCarturlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
