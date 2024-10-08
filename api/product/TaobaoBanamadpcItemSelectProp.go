package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBanamadpcItemSelectProp 获取子属性
// taobao.banamadpc.item.select.prop
//
// 巴拿马供应商通过此接口获取子属性
func TaobaoBanamadpcItemSelectProp(ctx context.Context, clt *core.SDKClient, req *product.TaobaoBanamadpcItemSelectPropAPIRequest, resp *product.TaobaoBanamadpcItemSelectPropAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
