package mtopopen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractMediaAudio 音频相关鉴权接口
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
func AlibabaInteractMediaAudio(ctx context.Context, clt *core.SDKClient, req *mtopopen.AlibabaInteractMediaAudioAPIRequest, resp *mtopopen.AlibabaInteractMediaAudioAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
