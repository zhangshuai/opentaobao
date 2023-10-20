package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorToast toast
// alibaba.interact.sensor.toast
//
// toast提示
func AlibabaInteractSensorToast(clt *core.SDKClient, req *interact.AlibabaInteractSensorToastAPIRequest, resp *interact.AlibabaInteractSensorToastAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
