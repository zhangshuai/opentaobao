package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDistributorsGet 获取分销商信息
// taobao.fenxiao.distributors.get
//
// 查询和当前登录供应商有合作关系的分销商的信息
func TaobaoFenxiaoDistributorsGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDistributorsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDistributorsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
