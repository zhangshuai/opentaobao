package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosflowWorkQueryvariables 获取指定流程上下文参数
// alibaba.mosflow.work.queryvariables
//
// 业务查询指定流程上下文内容
func AlibabaMosflowWorkQueryvariables(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosflowWorkQueryvariablesAPIRequest, resp *mos.AlibabaMosflowWorkQueryvariablesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
