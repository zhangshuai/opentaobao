package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslIteminfoBatchInsert 按商家批量写入商品接口
// taobao.uscesl.iteminfo.batch.insert
//
// 【电子价签】支持按照商家-门店维度批量写入商品数据
func TaobaoUsceslIteminfoBatchInsert(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslIteminfoBatchInsertAPIRequest, resp *uscesl.TaobaoUsceslIteminfoBatchInsertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
