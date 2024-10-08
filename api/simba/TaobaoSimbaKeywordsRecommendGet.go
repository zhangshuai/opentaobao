package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordsRecommendGet 取得一个推广组的推荐关键词列表
// taobao.simba.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
func TaobaoSimbaKeywordsRecommendGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordsRecommendGetAPIRequest, resp *simba.TaobaoSimbaKeywordsRecommendGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
