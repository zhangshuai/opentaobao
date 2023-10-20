package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenorderexceptionreport 订单异常通知接口
// taobao.qimen.orderexception.report
//
// WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
func Taobaoqimenorderexceptionreport(clt *core.SDKClient, req *qimen.TaobaoqimenorderexceptionreportAPIRequest, session string) (*qimen.TaobaoqimenorderexceptionreportAPIResponse, error) {
	var resp qimen.TaobaoqimenorderexceptionreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
