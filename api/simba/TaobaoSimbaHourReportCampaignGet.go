package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaHourReportCampaignGet 计划维度小时报表获取
// taobao.simba.hour.report.campaign.get
//
// 计划维度小时报表获取
func TaobaoSimbaHourReportCampaignGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaHourReportCampaignGetAPIRequest, resp *simba.TaobaoSimbaHourReportCampaignGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
