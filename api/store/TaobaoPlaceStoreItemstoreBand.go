package store

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoreItemstoreBand 门店商品关联绑定接口
// taobao.place.store.itemstore.band
//
// 商品和多个门店关系绑定接口
func TaobaoPlaceStoreItemstoreBand(ctx context.Context, clt *core.SDKClient, req *store.TaobaoPlaceStoreItemstoreBandAPIRequest, resp *store.TaobaoPlaceStoreItemstoreBandAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
