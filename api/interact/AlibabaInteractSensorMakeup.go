package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorMakeup 美妆虚拟试装
// alibaba.interact.sensor.makeup
//
// 手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
func AlibabaInteractSensorMakeup(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractSensorMakeupAPIRequest, resp *interact.AlibabaInteractSensorMakeupAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
