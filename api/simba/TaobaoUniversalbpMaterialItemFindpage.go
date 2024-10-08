package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpMaterialItemFindpage 分页查询商品信息
// taobao.universalbp.material.item.findpage
//
// 分页获取店铺内的商品列表
func TaobaoUniversalbpMaterialItemFindpage(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpMaterialItemFindpageAPIRequest, resp *simba.TaobaoUniversalbpMaterialItemFindpageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
