package scs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxMaterialMaterialFindpage 获取商品池
// taobao.onebp.dkx.material.material.findpage
//
// 获取商品池。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa。
func TaobaoOnebpDkxMaterialMaterialFindpage(ctx context.Context, clt *core.SDKClient, req *scs.TaobaoOnebpDkxMaterialMaterialFindpageAPIRequest, resp *scs.TaobaoOnebpDkxMaterialMaterialFindpageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
