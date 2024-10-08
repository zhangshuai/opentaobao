package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemPropertyDefQuery 交易猫商品属性定义查询
// alibaba.jym.item.property.def.query
//
// 查询商品发布属性定义详情
func AlibabaJymItemPropertyDefQuery(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemPropertyDefQueryAPIRequest, resp *product.AlibabaJymItemPropertyDefQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
