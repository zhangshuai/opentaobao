package traveltrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelBookinfoQuery 飞猪度假-订单二次预约查询接口
// alitrip.travel.bookinfo.query
//
// 飞猪度假订单二次预约详情查询接口
func AlitripTravelBookinfoQuery(ctx context.Context, clt *core.SDKClient, req *traveltrade.AlitripTravelBookinfoQueryAPIRequest, resp *traveltrade.AlitripTravelBookinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
