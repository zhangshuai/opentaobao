package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
排行榜歌曲获取 APIResponse
alibaba.xiami.api.rank.songs.get

获取歌曲排行榜
*/
type AlibabaXiamiApiRankSongsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiRankSongsGetResponse `json:"alibaba_xiami_api_rank_songs_get_response,omitempty"`
}

type AlibabaXiamiApiRankSongsGetResponse struct {

    // 返回结果
    Data   *RankSongsData `json:"data,omitempty"`

}