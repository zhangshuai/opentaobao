package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotDevicePost 商家提交设备信息
// alibaba.aliqin.fc.iot.device.post
//
// 物联网商家设备信息录入
func AlibabaAliqinFcIotDevicePost(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotDevicePostAPIRequest, resp *aliqin.AlibabaAliqinFcIotDevicePostAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
