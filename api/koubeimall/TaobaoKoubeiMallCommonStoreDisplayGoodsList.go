package koubeimall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStoreDisplayGoodsList 查询门店推荐菜信息
// taobao.koubei.mall.common.store.display.goods.list
//
// 提供查询口碑商圈内的门店推荐菜信息
func TaobaoKoubeiMallCommonStoreDisplayGoodsList(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonStoreDisplayGoodsListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
