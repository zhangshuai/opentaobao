package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinAxbVendorPushCallEvent 呼叫事件推送
// alibaba.aliqin.axb.vendor.push.call.event
//
// 呼叫事件推送
// 响铃时间、摘机事件
func AlibabaAliqinAxbVendorPushCallEvent(clt *core.SDKClient, req *alicom.AlibabaAliqinAxbVendorPushCallEventAPIRequest, resp *alicom.AlibabaAliqinAxbVendorPushCallEventAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
