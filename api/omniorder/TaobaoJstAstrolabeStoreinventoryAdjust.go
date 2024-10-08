package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryAdjust 后端商品库存占用调整接口
// taobao.jst.astrolabe.storeinventory.adjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
func TaobaoJstAstrolabeStoreinventoryAdjust(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryAdjustAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
