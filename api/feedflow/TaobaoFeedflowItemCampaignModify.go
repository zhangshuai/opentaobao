package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignModify 信息流修改计划
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
func TaobaoFeedflowItemCampaignModify(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignModifyAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
