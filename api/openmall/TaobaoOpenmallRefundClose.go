package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallRefundClose 关闭OpenMall退款单
// taobao.openmall.refund.close
//
// 关闭OpenMall退款单
func TaobaoOpenmallRefundClose(clt *core.SDKClient, req *openmall.TaobaoOpenmallRefundCloseAPIRequest, resp *openmall.TaobaoOpenmallRefundCloseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
