package alink

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkDeviceInfoUpdate 更新设备昵称等信息
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
func AlibabaAlinkDeviceInfoUpdate(ctx context.Context, clt *core.SDKClient, req *alink.AlibabaAlinkDeviceInfoUpdateAPIRequest, resp *alink.AlibabaAlinkDeviceInfoUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
