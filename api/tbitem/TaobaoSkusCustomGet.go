package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoSkusCustomGet 根据外部ID取商品SKU
// taobao.skus.custom.get
//
// 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku &lt;br/&gt;这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用)
func TaobaoSkusCustomGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoSkusCustomGetAPIRequest, resp *tbitem.TaobaoSkusCustomGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
