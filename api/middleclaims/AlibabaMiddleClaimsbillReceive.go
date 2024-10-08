package middleclaims

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/middleclaims"
)

// AlibabaMiddleClaimsbillReceive 国际化中台服务域接收理赔账单
// alibaba.middle.claimsbill.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔打款的处理后，将该打款结果返回至服务域
func AlibabaMiddleClaimsbillReceive(ctx context.Context, clt *core.SDKClient, req *middleclaims.AlibabaMiddleClaimsbillReceiveAPIRequest, resp *middleclaims.AlibabaMiddleClaimsbillReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
