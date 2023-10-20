package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSkuNewUpdate 更新sku销售属性标新状态
// tmall.item.sku.new.update
//
// 更新sku销售属性标新状态
func TmallItemSkuNewUpdate(clt *core.SDKClient, req *product.TmallItemSkuNewUpdateAPIRequest, resp *product.TmallItemSkuNewUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
