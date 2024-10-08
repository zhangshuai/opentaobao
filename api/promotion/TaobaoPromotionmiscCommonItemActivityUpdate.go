package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscCommonItemActivityUpdate 修改通用单品优惠活动
// taobao.promotionmisc.common.item.activity.update
//
// 修改通用单品优惠活动。
// 1、该接口只修改活动基本信息，如需要增加、删除参与该活动的商品及优惠，请调用taobao.promotionmisc.common.item.detail.add和taobao.promotionmisc.common.item.detail.delete接口
// 2、使用该接口时需要把未做修改的字段值也传入
func TaobaoPromotionmiscCommonItemActivityUpdate(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscCommonItemActivityUpdateAPIRequest, resp *promotion.TaobaoPromotionmiscCommonItemActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
