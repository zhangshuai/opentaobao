package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCrowdAdd 单品单元下，新增定向人群
// taobao.feedflow.item.crowd.add
//
// 单品单元下，新增定向人群
func TaobaoFeedflowItemCrowdAdd(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCrowdAddAPIRequest, resp *feedflow.TaobaoFeedflowItemCrowdAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
