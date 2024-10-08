package aliqin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIotModbind 物联网绑定/换绑API
// alibaba.aliqin.fc.iot.modbind
//
// 支持用户的设备的换绑和解绑操作
func AlibabaAliqinFcIotModbind(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotModbindAPIRequest, resp *aliqin.AlibabaAliqinFcIotModbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
