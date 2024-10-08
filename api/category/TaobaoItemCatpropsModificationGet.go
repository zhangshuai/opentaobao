package category

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// TaobaoItemCatpropsModificationGet 查询商品类目属性变更
// taobao.item.catprops.modification.get
//
// 查询商品类目属性变更信息
func TaobaoItemCatpropsModificationGet(ctx context.Context, clt *core.SDKClient, req *category.TaobaoItemCatpropsModificationGetAPIRequest, resp *category.TaobaoItemCatpropsModificationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
