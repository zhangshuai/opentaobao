package hotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotel"
)

// TaobaoXhotelInfoListGetForHello 哈罗获取酒店详情
// taobao.xhotel.info.list.get.for.hello
//
// 哈罗合作项目，供哈罗合作方批量和增量两种场景下查询已开通城市下的标准酒店及房型信息，不涉及用户登录信息
func TaobaoXhotelInfoListGetForHello(ctx context.Context, clt *core.SDKClient, req *hotel.TaobaoXhotelInfoListGetForHelloAPIRequest, resp *hotel.TaobaoXhotelInfoListGetForHelloAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
