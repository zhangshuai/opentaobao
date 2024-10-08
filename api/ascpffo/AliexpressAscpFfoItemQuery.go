package ascpffo

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

// AliexpressAscpFfoItemQuery AliExpress发货单明细分页查询API
// aliexpress.ascp.ffo.item.query
//
// AE履约发货单明细分页查询
func AliexpressAscpFfoItemQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpFfoItemQueryAPIRequest, resp *ascpffo.AliexpressAscpFfoItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
