package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderPackage 确认揽收商品
// alibaba.idle.rent.order.package
//
// 确认揽收
func AlibabaIdleRentOrderPackage(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderPackageAPIRequest, resp *idle.AlibabaIdleRentOrderPackageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
