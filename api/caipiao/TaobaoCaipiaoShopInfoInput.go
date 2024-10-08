package caipiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// TaobaoCaipiaoShopInfoInput 录入参加送彩票店铺信息
// taobao.caipiao.shop.info.input
//
// 录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。
func TaobaoCaipiaoShopInfoInput(ctx context.Context, clt *core.SDKClient, req *caipiao.TaobaoCaipiaoShopInfoInputAPIRequest, resp *caipiao.TaobaoCaipiaoShopInfoInputAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
