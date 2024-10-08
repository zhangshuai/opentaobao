package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderDtdQuery 门店自送根据核销码查订单
// taobao.omniorder.dtd.query
//
// 门店自送根据核销码码查询订单信息
func TaobaoOmniorderDtdQuery(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderDtdQueryAPIRequest, resp *omniorder.TaobaoOmniorderDtdQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
