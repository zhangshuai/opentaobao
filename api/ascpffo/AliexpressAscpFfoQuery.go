package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpFfoQuery AliExpress发货单查询API
// aliexpress.ascp.ffo.query
//
// AE 履约发货单分页查询接口
func AliexpressAscpFfoQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpFfoQueryAPIRequest, resp *ascpffo.AliexpressAscpFfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
