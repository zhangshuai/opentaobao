package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarKeywordsQscoreSplitGet (新)销量明星质量分相关接口
// taobao.simba.salestar.keywords.qscore.split.get
//
// 获取关键词新的质量分
func TaobaoSimbaSalestarKeywordsQscoreSplitGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest, resp *simba.TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
