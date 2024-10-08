package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// TaobaoAlitripTravelBaseinfoScenicsGet 【API3.0】基础信息获取接口：景点数据查询
// taobao.alitrip.travel.baseinfo.scenics.get
//
// 商品发布辅助接口，用于飞猪度假或门票商品发布时 获取可用的景点（及景点下收费项目）信息列表。
func TaobaoAlitripTravelBaseinfoScenicsGet(ctx context.Context, clt *core.SDKClient, req *travel.TaobaoAlitripTravelBaseinfoScenicsGetAPIRequest, resp *travel.TaobaoAlitripTravelBaseinfoScenicsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
