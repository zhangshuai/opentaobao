package middleclaims

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/middleclaims"
)

// AlibabaMiddleClaimsacceptReceive 国际化中台服务域接收保险公司理赔受理结果
// alibaba.middle.claimsaccept.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
func AlibabaMiddleClaimsacceptReceive(ctx context.Context, clt *core.SDKClient, req *middleclaims.AlibabaMiddleClaimsacceptReceiveAPIRequest, resp *middleclaims.AlibabaMiddleClaimsacceptReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
