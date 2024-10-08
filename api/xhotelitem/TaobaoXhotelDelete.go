package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelDelete 删除酒店接口
// taobao.xhotel.delete
//
// 删除飞猪酒店数据接口
func TaobaoXhotelDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
