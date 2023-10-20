package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// AlibabaAscpLogisticsInstantsonlineCalldelivery 同城配送在线下单正式下单呼叫运力
// alibaba.ascp.logistics.instantsonline.calldelivery
//
// 同城配送在线下单正式下单呼叫运力
func AlibabaAscpLogisticsInstantsonlineCalldelivery(clt *core.SDKClient, req *tblogistics.AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIRequest, resp *tblogistics.AlibabaAscpLogisticsInstantsonlineCalldeliveryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
