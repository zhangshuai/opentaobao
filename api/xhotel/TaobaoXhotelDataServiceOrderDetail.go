package xhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelDataServiceOrderDetail 服务订单详情
// taobao.xhotel.data.service.order.detail
//
// 服务订单详情top接口构建
func TaobaoXhotelDataServiceOrderDetail(ctx context.Context, clt *core.SDKClient, req *xhotel.TaobaoXhotelDataServiceOrderDetailAPIRequest, resp *xhotel.TaobaoXhotelDataServiceOrderDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
