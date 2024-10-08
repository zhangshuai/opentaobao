package yunosappstore

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunosappstore"
)

// YunosAppstoreOpenGetads 获取外投广告
// yunos.appstore.open.getads
//
// 将广告外投给外部合作伙伴
func YunosAppstoreOpenGetads(ctx context.Context, clt *core.SDKClient, req *yunosappstore.YunosAppstoreOpenGetadsAPIRequest, resp *yunosappstore.YunosAppstoreOpenGetadsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
