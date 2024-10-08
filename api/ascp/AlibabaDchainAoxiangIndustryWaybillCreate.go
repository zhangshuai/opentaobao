package ascp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangIndustryWaybillCreate 服务商开运单
// alibaba.dchain.aoxiang.industry.waybill.create
//
// 服务商开运单
func AlibabaDchainAoxiangIndustryWaybillCreate(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIRequest, resp *ascp.AlibabaDchainAoxiangIndustryWaybillCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
