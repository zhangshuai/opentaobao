package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TmallDeviceStoreFollowurlGet 获取店铺关注链接
// tmall.device.store.followurl.get
//
// 获取智能硬件上的关注店铺的URL
func TmallDeviceStoreFollowurlGet(ctx context.Context, clt *core.SDKClient, req *iot.TmallDeviceStoreFollowurlGetAPIRequest, resp *iot.TmallDeviceStoreFollowurlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
