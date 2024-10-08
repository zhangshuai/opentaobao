package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// TaobaoTmallgenieHotelwelcome 酒店欢迎词推送
// taobao.tmallgenie.hotelwelcome
//
// 推送欢迎词，让天猫精灵播放
func TaobaoTmallgenieHotelwelcome(ctx context.Context, clt *core.SDKClient, req *tmallgenie.TaobaoTmallgenieHotelwelcomeAPIRequest, resp *tmallgenie.TaobaoTmallgenieHotelwelcomeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
