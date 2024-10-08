package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivitySend 活动权益发放接口
// taobao.promotion.benefit.activity.send
//
// 活动权益发放接口，用于卖家针对活动进行权益发放
func TaobaoPromotionBenefitActivitySend(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivitySendAPIRequest, resp *promotion.TaobaoPromotionBenefitActivitySendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
