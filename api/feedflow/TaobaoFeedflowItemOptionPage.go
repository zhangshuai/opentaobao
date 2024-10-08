package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemOptionPage 分页查询定向标签列表
// taobao.feedflow.item.option.page
//
// 分页查询定向标签列表
func TaobaoFeedflowItemOptionPage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemOptionPageAPIRequest, resp *feedflow.TaobaoFeedflowItemOptionPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
