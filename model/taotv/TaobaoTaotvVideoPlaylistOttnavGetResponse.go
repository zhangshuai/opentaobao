package taotv

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ott播单 APIResponse
taobao.taotv.video.playlist.ottnav.get

根据聚焦播单ID拿到下面播单视频，根据左侧播单ID列表批量拿到播单信息
*/
type TaobaoTaotvVideoPlaylistOttnavGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTaotvVideoPlaylistOttnavGetResponse `json:"taobao_taotv_video_playlist_ottnav_get_response,omitempty"`
}

type TaobaoTaotvVideoPlaylistOttnavGetResponse struct {

    // result
    Result   *TaobaoTaotvVideoPlaylistOttnavGetResult `json:"result,omitempty"`

}