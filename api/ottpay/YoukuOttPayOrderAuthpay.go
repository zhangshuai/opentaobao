package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderAuthpay 委托代扣服务
// youku.ott.pay.order.authpay
//
// 应用中心sdk连续包月委托代扣服务
func YoukuOttPayOrderAuthpay(clt *core.SDKClient, req *ottpay.YoukuOttPayOrderAuthpayAPIRequest, resp *ottpay.YoukuOttPayOrderAuthpayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
