package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaKeywordidsChangedGet 获取修改的词ID
// taobao.simba.keywordids.changed.get
//
// 获取修改的词ID
func TaobaoSimbaKeywordidsChangedGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaKeywordidsChangedGetAPIRequest, resp *simba.TaobaoSimbaKeywordidsChangedGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
