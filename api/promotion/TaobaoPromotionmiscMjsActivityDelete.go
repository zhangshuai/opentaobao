package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscMjsActivityDelete 删除满就送活动
// taobao.promotionmisc.mjs.activity.delete
//
// 删除满就送活动
func TaobaoPromotionmiscMjsActivityDelete(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscMjsActivityDeleteAPIRequest, resp *promotion.TaobaoPromotionmiscMjsActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
