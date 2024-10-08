package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// AlibabaAilabsAligenieIotDeviceControlResult 设备控制结果
// alibaba.ailabs.aligenie.iot.device.control.result
//
// 智能IOT解决外部厂商在云云模式在用户通过天猫精灵下发设备指令过程中，厂商指令完成，回调结果通知
func AlibabaAilabsAligenieIotDeviceControlResult(ctx context.Context, clt *core.SDKClient, req *iot.AlibabaAilabsAligenieIotDeviceControlResultAPIRequest, resp *iot.AlibabaAilabsAligenieIotDeviceControlResultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
