package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaShopCategoryAllGet 全部店铺分类信息查询接口
// alibaba.shop.category.all.get
//
// 按照卖家身份查询全部分类信息
func AlibabaShopCategoryAllGet(ctx context.Context, clt *core.SDKClient, req *shop.AlibabaShopCategoryAllGetAPIRequest, resp *shop.AlibabaShopCategoryAllGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
