package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscMjsActivityListGet 查询满就送活动列表
// taobao.promotionmisc.mjs.activity.list.get
//
// 查询满就送活动列表。注意，该接口的返回值中，只包含活动的主要信息，如activity_id，name，description，start_time，end_time，type，participate_range。优惠的详细信息，请通过taobao.promotionmisc.mjs.activity.get获取。
func TaobaoPromotionmiscMjsActivityListGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscMjsActivityListGetAPIRequest, resp *promotion.TaobaoPromotionmiscMjsActivityListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
