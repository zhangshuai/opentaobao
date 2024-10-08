package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCampaignFindsubcampaignid 查询无界版计划对应的原场景计划id
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
func TaobaoUniversalbpCampaignFindsubcampaignid(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpCampaignFindsubcampaignidAPIRequest, resp *simba.TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
