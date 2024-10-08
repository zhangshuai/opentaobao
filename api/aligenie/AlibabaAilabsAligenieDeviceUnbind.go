package aligenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aligenie"
)

// AlibabaAilabsAligenieDeviceUnbind 设备解绑操作接口
// alibaba.ailabs.aligenie.device.unbind
//
// 让开发者能根据设备ID进行解绑操作的接口
func AlibabaAilabsAligenieDeviceUnbind(ctx context.Context, clt *core.SDKClient, req *aligenie.AlibabaAilabsAligenieDeviceUnbindAPIRequest, resp *aligenie.AlibabaAilabsAligenieDeviceUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
