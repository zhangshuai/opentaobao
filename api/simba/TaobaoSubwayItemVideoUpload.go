package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSubwayItemVideoUpload 创意视频上传
// taobao.subway.item.video.upload
//
// 为用户提供视频上传的功能
func TaobaoSubwayItemVideoUpload(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSubwayItemVideoUploadAPIRequest, resp *simba.TaobaoSubwayItemVideoUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
