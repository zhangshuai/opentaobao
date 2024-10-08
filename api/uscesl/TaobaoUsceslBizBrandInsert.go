package uscesl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/uscesl"
)

// TaobaoUsceslBizBrandInsert 新增电子价签商家
// taobao.uscesl.biz.brand.insert
//
// 一个电子价签业务身份下新增商家接口
func TaobaoUsceslBizBrandInsert(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslBizBrandInsertAPIRequest, resp *uscesl.TaobaoUsceslBizBrandInsertAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
