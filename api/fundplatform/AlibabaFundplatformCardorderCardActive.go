package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardorderCardActive 储值卡激活
// alibaba.fundplatform.cardorder.card.active
//
// 储值卡接货接口，可以通过外部订单号或者卡号进行批量激活。如果储值卡已经被激活过仍然幂等返回成功。资金平台保证批量激活时一定全部成功或全部失败。
// 如果批量激活储值卡时，如果部分储值卡处于已激活，部分储值卡处于未激活，则会返回失败
func AlibabaFundplatformCardorderCardActive(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderCardActiveAPIRequest, resp *fundplatform.AlibabaFundplatformCardorderCardActiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
