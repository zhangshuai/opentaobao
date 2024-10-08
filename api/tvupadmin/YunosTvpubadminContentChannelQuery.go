package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChannelQuery 迎客松影视频道查询
// yunos.tvpubadmin.content.channel.query
//
// 迎客松影视频道查询
func YunosTvpubadminContentChannelQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChannelQueryAPIRequest, resp *tvupadmin.YunosTvpubadminContentChannelQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
