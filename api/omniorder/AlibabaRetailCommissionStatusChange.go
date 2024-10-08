package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// AlibabaRetailCommissionStatusChange 分佣状态变更
// alibaba.retail.commission.status.change
//
// 分佣系统，分佣状态变更接口
func AlibabaRetailCommissionStatusChange(ctx context.Context, clt *core.SDKClient, req *omniorder.AlibabaRetailCommissionStatusChangeAPIRequest, resp *omniorder.AlibabaRetailCommissionStatusChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
