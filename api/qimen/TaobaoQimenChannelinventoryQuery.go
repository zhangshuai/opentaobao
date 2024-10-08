package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenChannelinventoryQuery 渠道库存查询接口
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
func TaobaoQimenChannelinventoryQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenChannelinventoryQueryAPIRequest, resp *qimen.TaobaoQimenChannelinventoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
