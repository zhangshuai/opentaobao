package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateUpdate 价格推送接口（全量更新）
// taobao.xhotel.rate.update
//
// 酒店产品库rate更新
func TaobaoXhotelRateUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
