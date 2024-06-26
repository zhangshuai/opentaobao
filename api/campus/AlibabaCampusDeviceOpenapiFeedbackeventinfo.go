package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusDeviceOpenapiFeedbackeventinfo IVS事件处理反馈接口
// alibaba.campus.device.openapi.feedbackeventinfo
//
// 提供给第三方ISV的的事件信息处理反馈的接口
func AlibabaCampusDeviceOpenapiFeedbackeventinfo(clt *core.SDKClient, req *campus.AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIRequest, resp *campus.AlibabaCampusDeviceOpenapiFeedbackeventinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
