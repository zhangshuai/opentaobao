package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScMaterialOptional 淘宝客-服务商-物料搜索
// taobao.tbk.sc.material.optional
//
// 服务商使用。支持入参推广者对应的“推广位”、关键词和相关筛选条件，获取对应的物料信息和推广者对应的推广链接。
func TaobaoTbkScMaterialOptional(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkScMaterialOptionalAPIRequest, resp *tbk.TaobaoTbkScMaterialOptionalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
