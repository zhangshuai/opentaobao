package ascpchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorPriceGet 链渠道中心淘外分销价格查询(分销商专用)
// alibaba.ascp.channel.distributor.price.get
//
// 此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
func AlibabaAscpChannelDistributorPriceGet(ctx context.Context, clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorPriceGetAPIRequest, resp *ascpchannel.AlibabaAscpChannelDistributorPriceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
