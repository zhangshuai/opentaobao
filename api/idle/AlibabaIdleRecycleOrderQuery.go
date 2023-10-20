package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderQuery 闲鱼回收订单查询V1
// alibaba.idle.recycle.order.query
//
// 查询回收订单
func AlibabaIdleRecycleOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderQueryAPIRequest, resp *idle.AlibabaIdleRecycleOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
