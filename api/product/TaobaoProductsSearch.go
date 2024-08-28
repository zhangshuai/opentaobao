package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoProductsSearch 搜索产品信息
// taobao.products.search
//
// 只有天猫商家发布商品时才需要用到，并非商品搜索api，当前暂不提供商品搜索api。&lt;br/&gt;二种方式搜索所有产品信息(二种至少传一种): &lt;br/&gt;
// 传入关键字q搜索&lt;br/&gt;
// 传入cid和props搜索&lt;br/&gt;
// 返回值支持:product_id,name,pic_path,cid,props,price,tsc&lt;br/&gt;
// 当用户指定了cid并且cid为垂直市场（3C电器城、鞋城）的类目id时，默认只返回小二确认&lt;br/&gt;的产品。如果用户没有指定cid，或cid为普通的类目，默认返回商家确认或小二确认的产&lt;br/&gt;品。如果用户自定了status字段，以指定的status类型为准。
// &lt;br/&gt;新增一：
//
//	传入suite_items_str 按规格搜索套装产品。
//	返回字段增加suite_items_str,is_suite_effecitve支持。
func TaobaoProductsSearch(ctx context.Context, clt *core.SDKClient, req *product.TaobaoProductsSearchAPIRequest, resp *product.TaobaoProductsSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
