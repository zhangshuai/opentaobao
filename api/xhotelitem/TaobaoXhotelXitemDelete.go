package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelXitemDelete 删除 x 元素
// taobao.xhotel.xitem.delete
//
// 删除 x 元素
func TaobaoXhotelXitemDelete(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelXitemDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelXitemDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
