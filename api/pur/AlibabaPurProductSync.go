package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurProductSync 同步产品
// alibaba.pur.product.sync
//
// 同步产品
func AlibabaPurProductSync(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurProductSyncAPIRequest, resp *pur.AlibabaPurProductSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
