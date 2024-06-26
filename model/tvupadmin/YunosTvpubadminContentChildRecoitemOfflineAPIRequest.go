package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildRecoitemOfflineAPIRequest 下线少儿推荐内容接口 API请求
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
type YunosTvpubadminContentChildRecoitemOfflineAPIRequest struct {
	model.Params
	// 推荐内容ID
	_recItemId int64
}

// NewYunosTvpubadminContentChildRecoitemOfflineRequest 初始化YunosTvpubadminContentChildRecoitemOfflineAPIRequest对象
func NewYunosTvpubadminContentChildRecoitemOfflineRequest() *YunosTvpubadminContentChildRecoitemOfflineAPIRequest {
	return &YunosTvpubadminContentChildRecoitemOfflineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChildRecoitemOfflineAPIRequest) Reset() {
	r._recItemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.recoitem.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRecItemId is RecItemId Setter
// 推荐内容ID
func (r *YunosTvpubadminContentChildRecoitemOfflineAPIRequest) SetRecItemId(_recItemId int64) error {
	r._recItemId = _recItemId
	r.Set("rec_item_id", _recItemId)
	return nil
}

// GetRecItemId RecItemId Getter
func (r YunosTvpubadminContentChildRecoitemOfflineAPIRequest) GetRecItemId() int64 {
	return r._recItemId
}

var poolYunosTvpubadminContentChildRecoitemOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChildRecoitemOfflineRequest()
	},
}

// GetYunosTvpubadminContentChildRecoitemOfflineRequest 从 sync.Pool 获取 YunosTvpubadminContentChildRecoitemOfflineAPIRequest
func GetYunosTvpubadminContentChildRecoitemOfflineAPIRequest() *YunosTvpubadminContentChildRecoitemOfflineAPIRequest {
	return poolYunosTvpubadminContentChildRecoitemOfflineAPIRequest.Get().(*YunosTvpubadminContentChildRecoitemOfflineAPIRequest)
}

// ReleaseYunosTvpubadminContentChildRecoitemOfflineAPIRequest 将 YunosTvpubadminContentChildRecoitemOfflineAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChildRecoitemOfflineAPIRequest(v *YunosTvpubadminContentChildRecoitemOfflineAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChildRecoitemOfflineAPIRequest.Put(v)
}
