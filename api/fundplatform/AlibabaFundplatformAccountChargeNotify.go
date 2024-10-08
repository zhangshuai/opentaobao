package fundplatform

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformAccountChargeNotify 账户充值成功通知
// alibaba.fundplatform.account.charge.notify
//
// 通知外部业务方充值成功
func AlibabaFundplatformAccountChargeNotify(ctx context.Context, clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountChargeNotifyAPIRequest, resp *fundplatform.AlibabaFundplatformAccountChargeNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
