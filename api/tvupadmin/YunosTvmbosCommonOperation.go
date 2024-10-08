package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvmbosCommonOperation 应用中心通用服务接口
// yunos.tvmbos.common.operation
//
// 应用中心相关接口的代理
func YunosTvmbosCommonOperation(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvmbosCommonOperationAPIRequest, resp *tvupadmin.YunosTvmbosCommonOperationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
