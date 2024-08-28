package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoProductAdd 上传一个产品，不包括产品非主图和属性图片
// taobao.product.add
//
// 获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取 &lt;br/&gt;传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid,&lt;br/&gt;调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.&lt;br/&gt;新增：套装产品发布，目前支持单件多个即 A*2 形式的套装
func TaobaoProductAdd(ctx context.Context, clt *core.SDKClient, req *product.TaobaoProductAddAPIRequest, resp *product.TaobaoProductAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
