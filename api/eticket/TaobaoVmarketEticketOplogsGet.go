package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketOplogsGet 电子凭证操作日志查询
// taobao.vmarket.eticket.oplogs.get
//
// 电子凭证核销日志查询
func TaobaoVmarketEticketOplogsGet(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketOplogsGetAPIRequest, resp *eticket.TaobaoVmarketEticketOplogsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
