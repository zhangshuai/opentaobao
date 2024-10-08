package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoUmpMbbsListGet 通过ids列表获取营销积木块列表
// taobao.ump.mbbs.list.get
//
// 通过营销积木id列表来获取营销积木块列表。
func TaobaoUmpMbbsListGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoUmpMbbsListGetAPIRequest, resp *promotion.TaobaoUmpMbbsListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
