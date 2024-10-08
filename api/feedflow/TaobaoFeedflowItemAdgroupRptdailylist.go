package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdgroupRptdailylist 推广单元分日数据查询
// taobao.feedflow.item.adgroup.rptdailylist
//
// 推广单元分日数据查询
func TaobaoFeedflowItemAdgroupRptdailylist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdgroupRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemAdgroupRptdailylistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
