package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdTargetTagListRecommendTag 给计划推荐标签
// alibaba.scbp.ad.target.tag.list.recommend.tag
//
// 给计划推荐标签
func AlibabaScbpAdTargetTagListRecommendTag(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagListRecommendTagAPIRequest, resp *scbp.AlibabaScbpAdTargetTagListRecommendTagAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
