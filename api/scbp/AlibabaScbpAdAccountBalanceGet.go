package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdAccountBalanceGet 查询账户余额
// alibaba.scbp.ad.account.balance.get
//
// 查询推广账户余额
func AlibabaScbpAdAccountBalanceGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdAccountBalanceGetAPIRequest, resp *scbp.AlibabaScbpAdAccountBalanceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
