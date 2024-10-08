package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeReviewIndexSync 新测评乐居指数接口
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
func AlibabaAlihouseNewhomeReviewIndexSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
