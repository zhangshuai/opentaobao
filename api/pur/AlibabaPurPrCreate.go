package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurPrCreate 下pr单
// alibaba.pur.pr.create
//
// 下pr单
func AlibabaPurPrCreate(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurPrCreateAPIRequest, resp *pur.AlibabaPurPrCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
