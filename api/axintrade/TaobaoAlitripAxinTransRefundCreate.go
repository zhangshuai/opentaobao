package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransRefundCreate 阿信创建退款单
// taobao.alitrip.axin.trans.refund.create
//
// 阿信供销平台-创建退款单服务
func TaobaoAlitripAxinTransRefundCreate(clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransRefundCreateAPIRequest, resp *axintrade.TaobaoAlitripAxinTransRefundCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
