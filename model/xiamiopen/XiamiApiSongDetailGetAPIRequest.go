package xiamiopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// XiamiApiSongDetailGetAPIRequest 获取歌曲详情 API请求
// xiami.api.song.detail.get
//
// 获取歌曲详情
type XiamiApiSongDetailGetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// NewXiamiApiSongDetailGetRequest 初始化XiamiApiSongDetailGetAPIRequest对象
func NewXiamiApiSongDetailGetRequest() *XiamiApiSongDetailGetAPIRequest {
	return &XiamiApiSongDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *XiamiApiSongDetailGetAPIRequest) Reset() {
	r._songIds = r._songIds[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiApiSongDetailGetAPIRequest) GetApiMethodName() string {
	return "xiami.api.song.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiApiSongDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamiApiSongDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSongIds is SongIds Setter
// 歌曲id
func (r *XiamiApiSongDetailGetAPIRequest) SetSongIds(_songIds []int64) error {
	r._songIds = _songIds
	r.Set("song_ids", _songIds)
	return nil
}

// GetSongIds SongIds Getter
func (r XiamiApiSongDetailGetAPIRequest) GetSongIds() []int64 {
	return r._songIds
}

var poolXiamiApiSongDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewXiamiApiSongDetailGetRequest()
	},
}

// GetXiamiApiSongDetailGetRequest 从 sync.Pool 获取 XiamiApiSongDetailGetAPIRequest
func GetXiamiApiSongDetailGetAPIRequest() *XiamiApiSongDetailGetAPIRequest {
	return poolXiamiApiSongDetailGetAPIRequest.Get().(*XiamiApiSongDetailGetAPIRequest)
}

// ReleaseXiamiApiSongDetailGetAPIRequest 将 XiamiApiSongDetailGetAPIRequest 放入 sync.Pool
func ReleaseXiamiApiSongDetailGetAPIRequest(v *XiamiApiSongDetailGetAPIRequest) {
	v.Reset()
	poolXiamiApiSongDetailGetAPIRequest.Put(v)
}
