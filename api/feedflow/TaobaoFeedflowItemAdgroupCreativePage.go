package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupCreativePage 信息流单元下查看创意
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
func TaobaoFeedflowItemAdgroupCreativePage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupCreativePageAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupCreativePageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
