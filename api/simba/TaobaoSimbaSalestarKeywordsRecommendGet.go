package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarKeywordsRecommendGet 销量明星api相关接口
// taobao.simba.salestar.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
func TaobaoSimbaSalestarKeywordsRecommendGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest, resp *simba.TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
